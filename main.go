package main

import "net/http"

func main() {
	cssHandler := http.FileServer(http.Dir("templates/css/"))
	imageHandler := http.FileServer(http.Dir("templates/images/"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/images/", http.StripPrefix("/images/", imageHandler))
	http.HandleFunc("/", indexHandler().Handle)
	http.HandleFunc("/about", aboutHandler().Handle)
	http.HandleFunc("/uploader", uploaderHandler().Handle)
	http.ListenAndServe(":8080", nil)
}

func indexHandler() *Handler {
	index := Page{Title: "index"}
	return &Handler{
		Page: index,
		Templates: []string{"templates/index.html", "templates/header.html"},
	}
}

func aboutHandler() *Handler {
	about := Page{Title: "about"}
	return &Handler{
		Page: about,
		Templates: []string{"templates/about.html", "templates/header.html"},
	}
}

func uploaderHandler() *Handler {
	uploader := Page{Title: "uploader"}
	return &Handler{
		Page: uploader,
		Templates: []string{"templates/uploader.html", "templates/header.html"},
	}
}
