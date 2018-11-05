package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/riona/blog/controller/article"
	"strconv"
	"context"
	"github.com/google/go-github/github"
	"os"
	"html/template"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../assets/templates/*.tmpl")

	//index page
	router.GET("/index", func(c *gin.Context) {

		finder, err := article.NewFinder()
		if err != nil {
			c.String(http.StatusBadRequest, "caused error new finder")
		}

		md := []byte("## hogehoge document")
		client := github.NewClient(nil)
		opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}
		body, _, err := client.Markdown(context.Background(), string(md), opt)

		t, err := template.New("../assets/templates/index.tmpl").Parse(body)
		if err != nil {
			c.String(http.StatusBadRequest, "caused error parse err: %v", err)
		}

		data := struct {
			Escaped    string
			NotEscaped template.HTML
		}{
			Escaped:    "<b>escaped<b>",
			NotEscaped: "<b>not escaped</b>",
		}

		err = t.Execute(os.Stdout, data)
		if err != nil {
			c.String(http.StatusBadRequest, "caused error execute err: %v", err)
		}

		articles := finder.FindArticleForIndex()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"articles": articles,
			"output":   body,
		})
	})

	// show page
	router.GET("/show/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "error parse string to int")
		}

		finder, err := article.NewFinder()
		if err != nil {
			c.String(http.StatusBadRequest, "caused error new finder")
		}

		article := finder.FindArticleByArticle(id)

		c.HTML(http.StatusOK, "show.tmpl", gin.H{
			"title": article.Title,
			"body":  article.Body,
			"time":  article.Date,
		})
	})

	router.Run(":8080")
}

type pageData struct {
	Body string
}
