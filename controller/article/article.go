package article

import (
	"github.com/gin-gonic/gin"
	"github.com/riona/blog/api"
	"net/http"
)

type IndexHandler struct{}

func GetIndex(c *gin.Context) {
	var err error
	api, err := api.NewFinder()
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	articles := api.FindArticleForIndex()
	for _, article := range articles {
		c.JSON(http.StatusOK, article)
	}
}
