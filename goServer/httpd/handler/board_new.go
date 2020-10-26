package handler

import (
	"mulTic/platform/mulTic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BoardNew(status *mulTic.Status) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := multicPostRequest{}
		c.Bind(&requestBody)

		n := requestBody.NextPlayer

		status.Restart(n)

		c.Status(http.StatusNoContent)
	}
}
