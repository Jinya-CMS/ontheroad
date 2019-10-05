package http

import (
	"html/template"
	"net/http"
)

func Render(tmpl string, tmplData interface{}, w http.ResponseWriter) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, tmplData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
