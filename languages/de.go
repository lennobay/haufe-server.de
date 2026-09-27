package languages

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var data_articles_index_de string = ""
var top_de string = "<!doctype html><html lang='de'><head><meta charset='UTF-8' /><meta name='viewport' content='width=device-width, initial-scale=1.0' /><title>Haufe Blog</title><link rel='stylesheet' href='/index.css' /><link rel='icon' type='image/png' href='/favicon-96x96.png' sizes='96x96'/><link rel='icon' type='image/svg+xml' href='/favicon.svg'/><link rel='shortcut icon' href='/favicon.ico' /><link rel='apple-touch-icon' sizes='180x180' href='/apple-touch-icon.png' /><link rel='manifest' href='/site.webmanifest' /><link rel='preconnect' href='https://fonts.googleapis.com'><link rel='preconnect' href='https://fonts.gstatic.com' crossorigin><link href='https://fonts.googleapis.com/css2?family=BBH+Bogle&display=swap' rel='stylesheet'></head><body><header><a href='/' ><b>Haufe Blog</b></a></header>"
var bottom_de string = "<footer><p>&copy Lennard Haufe</p><p>Gemacht mit GO</p><p>Keine Benutzung von KI</p><a href='/de/impressum_privatsphäre/'>Impressum & Privatsphäre</a></footer></body><script src='/index.js'></script></html>"
var tag_collection_de string = ""

func German_Sub_Page() {
	err := os.Mkdir("./website/de/", 0755)
	err = os.Mkdir("./website/de/artikel/", 0755)
	entries, err := os.ReadDir("./articles/de/")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles_de(e.Name())
	}
	log.Println(tag_collection_de)
	editArticlesPage_de()
	entries2, err := os.ReadDir("./other/de/")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries2 {
		readFileOthers_de(e.Name())
	}

}
func editArticlesPage_de() {

	err := os.WriteFile("./website/de/artikel/index.html", []byte(top_de+"<list-articles><div>"+data_articles_index_de+"</div></list-articles>"+bottom_de), 0755)
	if err != nil {
		panic(err)
	}
}
func analayzeFiles_de(filename string) {
	file, err := os.Open("./articles/de/" + filename)
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
				part_tags += "<a href='/lv/tags/" + part_2[i] + "'>" + part_2[i] + "</a>"
				tag_collection_de += "#" + part_2[i] + ":" + filename_new[0] + ":" + title_overview_articles + ":" + description_overview_articles + ":" + image_overview_articles
			}
			part_tags += "</articles-tags>"
			tags_page_article = part_tags
		} else if content == "" {
			end_text_page_article += "<br>"
		} else {
			end_text_page_article += "<p>" + content + "</p>"
		}

	}

	page_Article_de(filename_new[0], end_text_page_article, tags_page_article)
	page_Collection_Overview_Articles_de(filename_new[0], title_overview_articles, image_overview_articles, description_overview_articles, tags_page_article)

}
func page_Collection_Overview_Articles_de(filename string, title string, image string, description string, tags string) {
	data_articles_index_de += "<article-display><a href='/de/artikel/" + filename + ".html'>" + title + image + description + "</a></article-display>"
}
func page_Article_de(filename string, text string, tags string) {
	err := os.WriteFile("./website/de/artikel/"+filename+".html", []byte(top_de+"<articles-detail><articles-detail-down>"+text+"<br></articles-detail-down></articles-detail>"+bottom_de), 0755)
	if err != nil {
		panic(err)
	}
}

func readFileOthers_de(filename string) {
	file, err := os.Open("./other/de/" + filename)
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
	err = os.Mkdir("./website/de/"+filename_new[0], 0755)
	err = os.WriteFile("./website/de/"+filename_new[0]+"/index.html", []byte(top_de+"<articles-detail><div>"+new_text+"</div></articles-detail>"+bottom_de), 0755)
	if err != nil {
		panic(err)
	}
}
