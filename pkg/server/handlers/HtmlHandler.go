package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func HtmlHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fp := path.Join("static", "html", fmt.Sprintf("%s.html", name))
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
