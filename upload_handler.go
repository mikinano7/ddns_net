package main

import (
	"net/http"
	"io"
	"fmt"
	"os"
	"io/ioutil"
)

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
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("upload_file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		f, err := os.OpenFile("./files/uploader/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		http.Redirect(w, r, h.handler.Path, 301)
	}
}

func (h *UploadHandler) setContent() error {
	files , err := ioutil.ReadDir("./files/uploader/")
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
