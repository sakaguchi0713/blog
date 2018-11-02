package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/riona/blog/controller/article"
	"strconv"
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

		articles := finder.FindArticleForIndex()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"articles": articles,
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
			"body": article.Body,
			"time": article.Date,
		})
	})

	router.Run(":8080")
}
