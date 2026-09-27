package languages

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var data_articles_index_lv string = ""
var top_lv string = "<!doctype html><html lang='lv'><head><meta charset='UTF-8' /><meta name='viewport' content='width=device-width, initial-scale=1.0' /><title>Haufe Emuārs</title><link rel='stylesheet' href='/index.css' /><link rel='icon' type='image/png' href='/favicon-96x96.png' sizes='96x96'/><link rel='icon' type='image/svg+xml' href='/favicon.svg'/><link rel='shortcut icon' href='/favicon.ico' /><link rel='apple-touch-icon' sizes='180x180' href='/apple-touch-icon.png' /><link rel='manifest' href='/site.webmanifest' /><link rel='preconnect' href='https://fonts.googleapis.com'><link rel='preconnect' href='https://fonts.gstatic.com' crossorigin><link href='https://fonts.googleapis.com/css2?family=BBH+Bogle&display=swap' rel='stylesheet'></head><body><header><a href='/' ><b>Haufe Emuārs</b></a></header>"
var bottom_lv string = "<footer><p>&copy Lennards Haufe</p><p>Taisīta ar GO</p><p>Nekāda lietošana no MI</p><a href='/lv/juridiskā_informācija/'>Juridiskā informācija</a></footer></body><script src='/index.js'></script></html>"
var tag_collection_lv string = ""

func Latvian_Sub_Page() {
	err := os.Mkdir("./website/lv/", 0755)
	err = os.Mkdir("./website/lv/raksts/", 0755)
	entries, err := os.ReadDir("./articles/lv/")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles_lv(e.Name())
	}
	log.Println(tag_collection_lv)
	editArticlesPage_lv()
	entries2, err := os.ReadDir("./other/lv/")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries2 {
		readFileOthers_lv(e.Name())
	}

}
func editArticlesPage_lv() {

	err := os.WriteFile("./website/lv/raksts/index.html", []byte(top_lv+"<list-articles><div>"+data_articles_index_lv+"</div></list-articles>"+bottom_lv), 0755)
	if err != nil {
		panic(err)
	}
}
func analayzeFiles_lv(filename string) {
	file, err := os.Open("./articles/lv/" + filename)
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
				tag_collection_lv += "#" + part_2[i] + ":" + filename_new[0] + ":" + title_overview_articles + ":" + description_overview_articles + ":" + image_overview_articles
			}
			part_tags += "</articles-tags>"
			tags_page_article = part_tags
		} else if content == "" {
			end_text_page_article += "<br>"
		} else {
			end_text_page_article += "<p>" + content + "</p>"
		}

	}

	page_Article_lv(filename_new[0], end_text_page_article, tags_page_article)
	page_Collection_Overview_Articles_lv(filename_new[0], title_overview_articles, image_overview_articles, description_overview_articles, tags_page_article)

}
func page_Collection_Overview_Articles_lv(filename string, title string, image string, description string, tags string) {
	data_articles_index_lv += "<article-display><a href='/lv/raksts/" + filename + ".html'>" + title + image + description + "</a></article-display>"
}
func page_Article_lv(filename string, text string, tags string) {
	err := os.WriteFile("./website/lv/raksts/"+filename+".html", []byte(top_lv+"<articles-detail><articles-detail-down>"+text+"<br></articles-detail-down></articles-detail>"+bottom_lv), 0755)
	if err != nil {
		panic(err)
	}
}

func readFileOthers_lv(filename string) {
	file, err := os.Open("./other/lv/" + filename)
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
	err = os.Mkdir("./website/lv/"+filename_new[0], 0755)
	err = os.WriteFile("./website/lv/"+filename_new[0]+"/index.html", []byte(top_lv+"<articles-detail><div>"+new_text+"</div></articles-detail>"+bottom_lv), 0755)
	if err != nil {
		panic(err)
	}
}
