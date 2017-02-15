package main

import (
	"net/http"
	"io"
	"os"
	"io/ioutil"
)

const uploaderPath = "./files/uploader/"

type UploadHandler struct {
	handler *Handler
}

func (h *UploadHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := h.setContent()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			h.handler.Handle(w, r)
		}
	} else {
		files , err := ioutil.ReadDir(uploaderPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if len(files) > 99 {
			http.Error(w, "file upload limit exceeded.", http.StatusInternalServerError)
		}

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("upload_file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer file.Close()

		f, err := os.OpenFile(uploaderPath + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()

		io.Copy(f, file)
		http.Redirect(w, r, h.handler.Path, 301)
	}
}

func (h *UploadHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName != "" {
		err := os.Remove(uploaderPath + fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	http.Redirect(w, r, h.handler.Path, 301)
}

func (h *UploadHandler) setContent() error {
	files , err := ioutil.ReadDir(uploaderPath)
	if err != nil {
		return err
	}

	content := make([]string, len(files))
	for i, file := range files {
		content[i] = file.Name()
	}

	h.handler.Page = Page{
		h.handler.Page.Title,
		content,
	}
	return nil
}
