package view

import (
	"bytes"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"html/template"
	"net/http"
	"path/filepath"
)

// continue https://youtu.be/yNQ7_xz_UDk?t=6353

var (
	tpIndex      = parseTemplate("layout.tmpl", "index.tmpl")
	tpAdminLogin = parseTemplate("layout.tmpl", "admin/login.tmpl")
)

var m = minify.New()

const templateDir = "template"

func init() {
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

func joinTemplateDir(files ...string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(joinTemplateDir(files...)...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("layout")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buffer := bytes.Buffer{}
	err := t.Execute(&buffer, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.Minify("text/html", w, &buffer)
}

// Index render Index
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// AdminLogin render Admin Login
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}
