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

var (
	indexTemplate      = convertFileToTemplate("layout.tmpl", "index.tmpl")
	adminLoginTemplate = convertFileToTemplate("layout.tmpl", "admin/login.tmpl")
)

var templateMinifier = minify.New()

const defaultTemplateFolder = "template"

func init() {
	templateMinifier.AddFunc("text/html", html.Minify)
	templateMinifier.AddFunc("text/css", css.Minify)
	templateMinifier.AddFunc("text/javascript", js.Minify)
}

func joinFolderName(files ...string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(defaultTemplateFolder, f)
	}
	return r
}

func convertFileToTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(joinFolderName(files...)...)
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
	templateMinifier.Minify("text/html", w, &buffer)
}

// RenderIndexPage render index page
func RenderIndexPage(w http.ResponseWriter, data interface{}) {
	render(indexTemplate, w, data)
}

// RenderAdminLoginPage render admin login page
func RenderAdminLoginPage(w http.ResponseWriter, data interface{}) {
	render(adminLoginTemplate, w, data)
}
