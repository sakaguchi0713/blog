package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riona/blog/controller/article"
	"net/http"
	"fmt"
)

func main() {
	router := gin.New()

	//index page
	router.GET("/index", func(c *gin.Context) {
		finder, err := article.NewFinder()
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		articles := finder.FindArticleForIndex()
		for key, values := range articles {
			for _, value := range values {
				str := fmt.Sprint(key)
				c.String(http.StatusOK, "ID: %s, \nTITLE: %s \n", str, value.Title)
			}
		}
	})
	//router.Run(":8080")
	//
	//router.GET("/show", func(c *gin.Context) {
	//	finder, err := article.NewFinder()
	//	if err != nil {
	//		c.Status(http.StatusBadRequest)
	//	}
	//
	//	articles := finder.FindArticleForIndex()
	//	for key, values := range articles {
	//		for _, value := range values {
	//			c.String(http.StatusOK, "ID: %d, \nTITLE: %s \n", key, value)
	//		}
	//	}
	//})
	//router.Run(":8080")
}

