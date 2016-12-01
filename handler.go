package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title     string
	Content   interface{}
}

type Handler struct {
	Page     Page
	Templates []string
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles(h.Templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, h.Page);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
