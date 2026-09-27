package main

import (
	"log"
	"net/http"

	"haufe-server.de/languages"
)

func main() {
	languages.English_Sub_Page()
	languages.Latvian_Sub_Page()
	languages.German_Sub_Page()
	fileserver := http.FileServer(http.Dir("./website"))

	http.Handle("/", fileserver)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}

}
