package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"github.com/jinzhu/gorm"
)

type Page struct {
	Title string
	Count int
}

type Article struct {
	Title string
	Body  string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello, World.", 2}
	tmpl := template.Must(template.ParseFiles("layout.html"))

	err := tmpl.ExecuteTemplate(w, "layout.html", page)
	if err != nil {
		log.Fatal(err)
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	article := Article{
		Title: "Hello!",
		Body:  "World",
	}
	fmt.Fprint(w, article)
}

func connection () *gorm.DB{
	db, err := gorm.Open("mysql", "root:password@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("error connection mysql. err: %v", err)
	}
	return db
}

func main() {
	http.HandleFunc("/create", createArticle)
	http.ListenAndServe(":8080", nil)
}
