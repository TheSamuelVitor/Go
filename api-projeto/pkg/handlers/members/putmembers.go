package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutMembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "putMembers",
	})
}