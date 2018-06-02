package main

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Top struct {
	Title string
	Body  string
}

type Article struct {
	Title string
	Body  string
	Date  time.Time
}

//use html sample
//func viewHandler(w http.ResponseWriter, r *http.Request) {
//page := Page{"Hello, World.", 2}
//tmpl := template.Must(template.ParseFiles("layout.html"))
//
//err := tmpl.ExecuteTemplate(w, "layout.html", page)
//if err != nil {
//	log.Fatal(err)
//}
//}

func show(w http.ResponseWriter, r *http.Request) {
	article := Article{
		"Hello!",
		"World",
		time.Now(),
	}
	tmpl := template.Must(template.ParseFiles("layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout.html", article)
	if err != nil {
		log.Fatal(err)
	}
}

func getTop(w http.ResponseWriter, r *http.Request) {
	page := Top{Title: "MyBlog", Body: "This is my Blur"}
	fmt.Fprint(w, page)
}

type LineOfLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}

var TemplateOfLog = `
Remote address:   {{.RemoteAddr}}
Content-Type:     {{.ContentType}}
HTTP method:      {{.Method}}

path:
{{.Path}}

query string:
{{.Query}}

body:             
{{.Body}}

	`

func Log(handler http.Handler) http.Handler {
	//handlerをHandlerFuncにします
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//bufferを作るよ
		bufbody := new(bytes.Buffer)
		//requestのbodyを読み込んでる的な
		bufbody.ReadFrom(r.Body)
		//読み込んだやつをストリングにしてる時な
		body := bufbody.String()

		line := LineOfLog{
			//web要求を行ったクライアントのIPアドレス
			r.RemoteAddr,
			//content-typeのheaderを読み込む
			r.Header.Get("Content-Type"),
			//要求されたurl
			r.URL.Path,
			//生クエリ保存(なんだかわからん)
			r.URL.RawQuery,
			//要求されたアクション
			r.Method, body,
		}
		tmpl, err := template.New("line").Parse(TemplateOfLog)
		if err != nil {
			panic(err)
		}

		bufline := new(bytes.Buffer)
		//bufferにlineを書き込む
		err = tmpl.Execute(bufline, line)
		if err != nil {
			panic(err)
		}
		//logでbufferを読み込んでストリングで返す
		log.Printf(bufline.String())
		//hundlerfunc返します
		handler.ServeHTTP(w, r)
	})
}

type mysql struct {
	adapter  string `yaml:"adapter"`
	encoding string `yaml:"encording"`
	database string `yml:"database"`
	pool     string `yaml:"pool"`
	host     string `yaml:"host"`
	port     string `yaml:"port"`
	user     string `yaml:"user"`
	password string `yaml:"password"`
}

func migration(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	var m mysql
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", m.user, m.host, m.port, m.database))
}

func rooter() {
	http.HandleFunc("/show", show)
	http.HandleFunc("/", getTop)
}

func main() {
	rooter()
	log.Print("Started server.")
	err := http.ListenAndServe(":8080", Log(http.DefaultServeMux))
	if err != nil {
		panic(err)
		log.Fatalf("failed rooter err: %v", err)
	}
}
