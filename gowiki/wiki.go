package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mwl/gowiki/config"
	"net/http"
	"regexp"
)

// Page ..
type Page struct {
	Title string
	Body  []byte
}

var configInfo *config.ConfigInfo

func init() {
	log.SetPrefix("wiki: ")

}

func (p *Page) save() error {
	filename := configInfo.DataPath + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Page, error) {
	filename := configInfo.DataPath + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	templates := template.Must(template.ParseFiles(configInfo.TmplPath+"edit.html", configInfo.TmplPath+"view.html"))

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func viewHandle(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		// return
		p = &Page{Title: title}
	}
	renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("无效 Page Title")
	}
	return m[2], nil
}
func makeHandle(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
func main() {
	configInfo = config.NewConfig("config.yml")
	fmt.Println(configInfo)
	http.HandleFunc("/view/", makeHandle(viewHandle))
	http.HandleFunc("/edit/", makeHandle(editHandler))
	http.HandleFunc("/save/", makeHandle(saveHandler))
	log.Fatal(http.ListenAndServe(":"+configInfo.Port, nil))
}
