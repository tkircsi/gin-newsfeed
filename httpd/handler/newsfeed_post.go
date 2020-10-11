package handler

import (
	"net/http"

	"github.com/tkircsi/gin-newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// NewsfeedPost function
func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request",
			})
			return
		}
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)
		c.JSON(http.StatusCreated, item)
	}
}
