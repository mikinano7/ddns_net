package main

import "net/http"

func main() {
	cssHandler := http.FileServer(http.Dir("templates/css/"))
	imageHandler := http.FileServer(http.Dir("templates/images/"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/images/", http.StripPrefix("/images/", imageHandler))
	http.HandleFunc("/", indexHandler("/").Handle)
	http.HandleFunc("/about", aboutHandler("/about").Handle)
	http.HandleFunc("/uploader", uploaderHandler("/uploader").Handle)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(path string) *Handler {
	index := Page{Title: "index"}
	return &Handler{
		Page: index,
		Path: path,
		Templates: []string{"templates/index.html", "templates/header.html"},
	}
}

func aboutHandler(path string) *Handler {
	about := Page{Title: "about"}
	return &Handler{
		Page: about,
		Path: path,
		Templates: []string{"templates/about.html", "templates/header.html"},
	}
}

func uploaderHandler(path string) *Handler {
	uploader := Page{Title: "uploader"}
	return &Handler{
		Page: uploader,
		Path: path,
		Templates: []string{"templates/uploader.html", "templates/header.html"},
	}
}
