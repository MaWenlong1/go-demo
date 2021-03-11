package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mwl/gowiki/config"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
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
	// log.Fatal(http.ListenAndServe(":"+configInfo.Port, nil))
	// enum1()
	test1()
	// runtime.GOMAXPROCS(NCPU)
	fmt.Println("cpu:" + strconv.Itoa(runtime.NumCPU()))
	// 从无缓冲信道进行的接收，要发生在对该信道进行的发送完成之前。
	go f()
	c <- 0
	print(a)

}

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}

var slice []func()

func test1() {
	sli := []int{1, 2, 3, 4, 5}
	// v 都是只创建一次，然后循环中赋值。
	for _, v := range sli {
		fmt.Println(&v)
		// v := v //解决方法
		slice = append(slice, func() {
			fmt.Println(v * v) // 直接打印结果
		})
	}

	for _, val := range slice {
		val()
	}
}

func enum1() {
	// ByteSize 。。
	type ByteSize float64
	const (
		// 通过赋予空白标识符来忽略第一个值
		_  ByteSize = iota // ignore first value by assigning to blank identifier
		KB          = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Println(KB)
	fmt.Println(home, user, gopath)
	fmt.Println(os.Hostname())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getgid())
	fmt.Println(os.Getpid())
}
