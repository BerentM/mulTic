package handler

import (
	"mulTic/platform/mulTic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type multicPostRequest struct {
	Board string `json:"board"`
}

func BoardPost(status *mulTic.Status) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := multicPostRequest{}
		c.Bind(&requestBody)

		s := requestBody.Board

		status.Update(s)

		c.Status(http.StatusNoContent)
	}
}
