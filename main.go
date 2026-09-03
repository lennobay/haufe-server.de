package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
)

var data_articles_index string = ""
var top string = "<!doctype html><html lang='en'><head><meta charset='UTF-8' /><meta name='viewport' content='width=device-width, initial-scale=1.0' /><title>Haufenet</title><link rel='stylesheet' href='/index.css' /><link rel='icon' type='image/png' href='/favicon-96x96.png' sizes='96x96'/><link rel='icon' type='image/svg+xml' href='/favicon.svg'/><link rel='shortcut icon' href='/favicon.ico' /><link rel='apple-touch-icon' sizes='180x180' href='/apple-touch-icon.png' /><link rel='manifest' href='/site.webmanifest' /><link rel='preconnect' href='https://fonts.googleapis.com'><link rel='preconnect' href='https://fonts.gstatic.com' crossorigin><link href='https://fonts.googleapis.com/css2?family=BBH+Bogle&display=swap' rel='stylesheet'></head><body><header><a href='/' ><b>Haufenet</b></a></header>"
var bottom string = "<footer><p>&copy Lennard Haufe<p><p>No usage of AI</p><a href='/imprint_privacy/'>Imprint and Privacy</a></footer></body><script src='/index.js'></script></html>"

func main() {
	entries, err := os.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {

		ReadFileArticles(e.Name())
		SetupDataIndexArticles(e.Name())
	}
	log.Print("Server running on port 8090")
	EditArticlesPage()
	ReadFileOthers("imprint_privacy.md")
	fileserver := http.FileServer(http.Dir("./website"))
	http.Handle("/", fileserver)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}

}
func EditArticlesPage() {

	err := os.WriteFile("./website/articles/index.html", []byte(top+"<list-articles><h1 id='top-article-list'>Articles</h1><div>"+data_articles_index+"</div></list-articles>"+bottom), 0755)
	if err != nil {
		panic(err)
	}
}
func SetupDataIndexArticles(filename string) {

	file, err := os.Open("./articles/" + filename)
	if err != nil {
		log.Fatalf("not opening")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var title string = ""
	var description string = ""
	var tags string = ""
	var image string = ""
	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, "###") {
			continue
		} else if strings.Contains(content, "##") {
			continue
		} else if strings.Contains(content, "#") {
			part := strings.Split(content, `#`)
			title = "<h1>" + part[1] + "</h1>"
		} else if strings.Contains(content, "DESCRIPTION:") {
			part := strings.Split(content, `DESCRIPTION:`)
			description = "<p>" + part[1] + "</p>"
		} else if strings.Contains(content, "HEADIMAGE:") {
			part := strings.Split(content, `HEADIMAGE:`)
			image = "<img src='/pictures/" + part[1] + "'>"
		} else if strings.Contains(content, "TAGS:") {
			part := strings.Split(content, `TAGS:`)
			part_2 := strings.Split(part[1], ",")
			part_tags := "<articles-tags>"
			for i := 0; i < len(part_2); i++ {
				part_tags += "<a href='/tags/" + part_2[i] + "'>" + part_2[i] + "</a>"
			}
			part_tags += "</articles-tags>"
			tags = part_tags
		} else {
			continue
		}
	}

	filename_new := strings.Split(filename, ".md")
	data_articles_index += "<article-display><a href='/articles/" + filename_new[0] + ".html'>" + title + image + description + "</a>" + tags + "</article-display>"

}

func ReadFileArticles(filename string) {
	file, err := os.Open("./articles/" + filename)
	if err != nil {
		log.Fatalf("not opening")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var new_text string = ""
	var tags string = ""
	for scanner.Scan() {
		content := scanner.Text()
		var end_type string = ""

		if strings.Contains(content, "###") {
			part := strings.Split(content, `###`)
			end_type = "<h3>" + part[1] + "</h3>"
		} else if strings.Contains(content, "##") {
			part := strings.Split(content, `##`)
			end_type = "<h2>" + part[1] + "</h2>"
		} else if strings.Contains(content, "#") {
			part := strings.Split(content, `#`)
			end_type = "<h1>" + part[1] + "</h1>"
		} else if strings.Contains(content, "HEADIMAGE:") {
			part := strings.Split(content, `HEADIMAGE:`)
			end_type = "<img src='/pictures/" + part[1] + "'>"
		} else if strings.Contains(content, "[img]:") {
			part := strings.Split(content, `[img]:`)
			end_type = "<img src='/pictures/" + part[1] + "'>"
		} else if strings.Contains(content, "DESCRIPTION:") {
			end_type = ""
		} else if strings.Contains(content, "TAGS:") {
			part := strings.Split(content, `TAGS:`)
			part_2 := strings.Split(part[1], ",")
			part_tags := "<articles-tags>"
			for i := 0; i < len(part_2); i++ {
				part_tags += "<a href='/tags/" + part_2[i] + "'>" + part_2[i] + "</a>"
			}
			part_tags += "</articles-tags>"
			tags = part_tags
		} else if content == "" {
			end_type = "<br>"
		} else {
			end_type = "<p>" + content + "</p>"
		}
		new_text += end_type

	}
	filename_new := strings.Split(filename, ".md")
	err = os.WriteFile("./website/articles/"+filename_new[0]+".html", []byte(top+"<articles-detail><articles-detail-down>"+new_text+"<h4>Tags</h4>"+tags+"<br></articles-detail-down></articles-detail>"+bottom), 0755)
	if err != nil {
		panic(err)
	}
}

func ReadFileOthers(filename string) {
	file, err := os.Open("./other/" + filename)
	if err != nil {
		log.Fatalf("not opening")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var new_text string = ""
	for scanner.Scan() {
		content := scanner.Text()
		var end_type string = ""

		if strings.Contains(content, "###") {
			part := strings.Split(content, `###`)
			end_type = "<h3>" + part[1] + "</h3>"
		} else if strings.Contains(content, "##") {
			part := strings.Split(content, `##`)
			end_type = "<h2>" + part[1] + "</h2>"
		} else if strings.Contains(content, "#") {
			part := strings.Split(content, `#`)
			end_type = "<h1>" + part[1] + "</h1>"
		} else if content == "" {
			end_type = "<br>"
		} else {
			end_type = "<p>" + content + "</p>"
		}
		new_text += end_type

	}
	filename_new := strings.Split(filename, ".md")
	err = os.Mkdir("./website/"+filename_new[0], 0755)
	err = os.WriteFile("./website/"+filename_new[0]+"/index.html", []byte(top+"<articles-detail><div>"+new_text+"</div></articles-detail>"+bottom), 0755)
	if err != nil {
		panic(err)
	}
}
