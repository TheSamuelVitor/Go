package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutTask(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK, gin.H{
			"message": "put task",
		})
}