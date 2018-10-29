package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router := gin.Default()

	router.GET("blogs", GetBlogs)
	r.Run() // listen and server on 0.0.0.0:8080
}
