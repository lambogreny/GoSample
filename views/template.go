package views

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/yuttasakcom/GoAPI/models"
)

var (
	tpIndex           = parseTemplate("root.tmpl", "index.tmpl")
	tpNewsID          = parseTemplate("root.tmpl", "newsid.tmpl")
	tpAdminLogin      = parseTemplate("root.tmpl", "admin/login.tmpl")
	tpAdminList       = parseTemplate("root.tmpl", "admin/list.tmpl")
	tpAdminCreateNews = parseTemplate("root.tmpl", "admin/create.tmpl")
	tpAdminEditNews   = parseTemplate("root.tmpl", "admin/edit.tmpl")
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
	t = t.Lookup("root")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.Minify("text/html", w, &buf)
}

type IndexData struct {
	List []*models.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// NewsID renders newsid view
func NewsID(w http.ResponseWriter, data interface{}) {
	render(tpNewsID, w, data)
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList renders admin login view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminCreateNews renders admin create form
func AdminCreateNews(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreateNews, w, data)
}

// AdminEditNews renders admin create form
func AdminEditNews(w http.ResponseWriter, data interface{}) {
	render(tpAdminEditNews, w, data)
}
