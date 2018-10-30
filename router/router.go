package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riona/blog/controller/article"
)

func main() {
	router := gin.New()

	router.GET("/index", article.GetIndex)
	router.Run(":8080")
}
