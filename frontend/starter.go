package main

import (
	"net/http"
	"crawler/frontend/controller"
)

func main() {
	// serve css and js files
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))

	// handle /search
	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)

	if err != nil{
		panic(err)
	}
}
