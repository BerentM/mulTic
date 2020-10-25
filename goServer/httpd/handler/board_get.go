package handler

import (
	"mulTic/platform/mulTic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BoardGet(status *mulTic.Status) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := status.Show()
		c.JSON(http.StatusOK, results)
	}
}
