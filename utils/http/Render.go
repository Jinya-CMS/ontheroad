package http

import (
	"github.com/russross/blackfriday"
	"go.jinya.de/ontheroad/database"
	"html/template"
	"net/http"
	"strings"
	"time"
)

func RenderSingle(tmpl string, tmplData interface{}, w http.ResponseWriter) {
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
}

func RenderAdmin(tmpl string, tmplData interface{}, w http.ResponseWriter) {
	layout, err := template.New("layout").Funcs(template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
	}).ParseFiles(tmpl, "templates/admin/layout.html.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = layout.ExecuteTemplate(w, "layout", tmplData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func RenderFrontend(tmpl string, tmplData interface{}, r *http.Request, w http.ResponseWriter) {
	layout, err := template.New("layout").Funcs(template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"markdown": func(input string) template.HTML {
			return template.HTML(blackfriday.Run([]byte(input)))
		},
		"urlStartsWith": func(part string) bool {
			return strings.HasPrefix(r.URL.Path, part)
		},
		"urlEndsWith": func(part string) bool {
			return strings.HasSuffix(r.URL.Path, part)
		},
		"formatDate": func(date time.Time) string {
			return date.Format("02.01.2006")
		},
		"getConfig": func(key string) string {
			config, err := database.GetConfiguration(key)
			if err != nil {
				return ""
			}

			return config.Value
		},
	}).ParseFiles(tmpl, "templates/frontend/layout.html.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = layout.ExecuteTemplate(w, "layout", tmplData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
