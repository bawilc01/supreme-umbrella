//get endpoint
package handler

import (
	"net/http"

	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed *newsfeed.List) gin.HandlerFunc {
	return func (c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"hello": "Found me",
		})
	}
}