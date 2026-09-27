package languages

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var data_articles_index string = ""
var top string = "<!doctype html><html lang='en'><head><meta charset='UTF-8' /><meta name='viewport' content='width=device-width, initial-scale=1.0' /><title>Haufe Blog</title><link rel='stylesheet' href='/index.css' /><link rel='icon' type='image/png' href='/favicon-96x96.png' sizes='96x96'/><link rel='icon' type='image/svg+xml' href='/favicon.svg'/><link rel='shortcut icon' href='/favicon.ico' /><link rel='apple-touch-icon' sizes='180x180' href='/apple-touch-icon.png' /><link rel='manifest' href='/site.webmanifest' /><link rel='preconnect' href='https://fonts.googleapis.com'><link rel='preconnect' href='https://fonts.gstatic.com' crossorigin><link href='https://fonts.googleapis.com/css2?family=BBH+Bogle&display=swap' rel='stylesheet'></head><body><header><a href='/' ><b>Haufe Blog</b></a></header>"
var bottom string = "<footer><p>&copy Lennard Haufe</p><p>Compiled with GO</p><p>No usage of AI</p><a href='/en/imprint_privacy/'>Imprint and Privacy</a></footer></body><script src='/index.js'></script></html>"
var tag_collection string = ""

func English_Sub_Page() {
	err := os.Mkdir("./website/en/", 0755)
	err = os.Mkdir("./website/en/articles/", 0755)
	entries, err := os.ReadDir("./articles/en/")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles(e.Name())
	}
	log.Println(tag_collection)
	editArticlesPage()
	readFileOthers("imprint_privacy.md")

}
func editArticlesPage() {

	err := os.WriteFile("./website/en/articles/index.html", []byte(top+"<list-articles><div>"+data_articles_index+"</div></list-articles>"+bottom), 0755)
	if err != nil {
		panic(err)
	}
}
func analayzeFiles(filename string) {
	file, err := os.Open("./articles/en/" + filename)
	if err != nil {
		log.Fatalf("not opening")
	}
	filename_new := strings.Split(filename, ".md")

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var end_text_page_article string = ""
	var tags_page_article string = ""
	var title_overview_articles string = ""
	var description_overview_articles string = ""
	var image_overview_articles string = ""
	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, "###") {
			part := strings.Split(content, `###`)
			end_text_page_article += "<h3>" + part[1] + "</h3>"
		} else if strings.Contains(content, "##") {
			part := strings.Split(content, `##`)
			end_text_page_article += "<h2>" + part[1] + "</h2>"
		} else if strings.Contains(content, "#") {
			part := strings.Split(content, `#`)
			end_text_page_article += "<h1>" + part[1] + "</h1>"
			title_overview_articles = "<h1>" + part[1] + "</h1>"
		} else if strings.Contains(content, "HEADIMAGE:") {
			part := strings.Split(content, `HEADIMAGE:`)
			end_text_page_article += "<img src='/pictures/" + part[1] + "'>"
			image_overview_articles = "<img src='/pictures/" + part[1] + "'>"
		} else if strings.Contains(content, "DESCRIPTION:") {
			part := strings.Split(content, `DESCRIPTION:`)
			description_overview_articles = "<p>" + part[1] + "</p>"
		} else if strings.Contains(content, "[img]:") {
			part := strings.Split(content, `[img]:`)
			end_text_page_article += "<img src='/pictures/" + part[1] + "'>"
		} else if strings.Contains(content, "DESCRIPTION:") {

			end_text_page_article = ""
		} else if strings.Contains(content, "DATE:") {
			part := strings.Split(content, "DATE:")
			end_text_page_article += "<i>" + part[1] + "</i><br>"
		} else if strings.Contains(content, "TAGS:") {
			part := strings.Split(content, `TAGS:`)
			part_2 := strings.Split(part[1], ",")
			part_tags := "<articles-tags>"
			for i := 0; i < len(part_2); i++ {
				part_tags += "<a href='/en/tags/" + part_2[i] + "'>" + part_2[i] + "</a>"
				tag_collection += "#" + part_2[i] + ":" + filename_new[0] + ":" + title_overview_articles + ":" + description_overview_articles + ":" + image_overview_articles
			}
			part_tags += "</articles-tags>"
			tags_page_article = part_tags
		} else if content == "" {
			end_text_page_article += "<br>"
		} else {
			end_text_page_article += "<p>" + content + "</p>"
		}

	}

	page_Article(filename_new[0], end_text_page_article, tags_page_article)
	page_Collection_Overview_Articles(filename_new[0], title_overview_articles, image_overview_articles, description_overview_articles, tags_page_article)

}
func page_Collection_Overview_Articles(filename string, title string, image string, description string, tags string) {
	data_articles_index += "<article-display><a href='/en/articles/" + filename + ".html'>" + title + image + description + "</a></article-display>"
}
func page_Article(filename string, text string, tags string) {
	err := os.WriteFile("./website/en/articles/"+filename+".html", []byte(top+"<articles-detail><articles-detail-down>"+text+"<br></articles-detail-down></articles-detail>"+bottom), 0755)
	if err != nil {
		panic(err)
	}
}

func readFileOthers(filename string) {
	file, err := os.Open("./other/en/" + filename)
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
	err = os.Mkdir("./website/en/"+filename_new[0], 0755)
	err = os.WriteFile("./website/en/"+filename_new[0]+"/index.html", []byte(top+"<articles-detail><div>"+new_text+"</div></articles-detail>"+bottom), 0755)
	if err != nil {
		panic(err)
	}
}
