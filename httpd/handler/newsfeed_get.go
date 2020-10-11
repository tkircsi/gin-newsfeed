package handler

import (
	"net/http"

	"github.com/tkircsi/gin-newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

// NewsfeedGet function
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
