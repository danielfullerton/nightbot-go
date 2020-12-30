package handlers

import (
	"html/template"
	"net/http"
	"path"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("dist", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
