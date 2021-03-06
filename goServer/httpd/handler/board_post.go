package handler

import (
	"mulTic/platform/mulTic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type multicPostRequest struct {
	NextPlayer string `json:"nextPlayer"`
	Board      string `json:"board"`
}

func BoardPost(status *mulTic.Status) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := multicPostRequest{}
		c.Bind(&requestBody)

		b := requestBody.Board

		status.Update(b)

		c.Status(http.StatusNoContent)
	}
}
