package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "put Projects",
	})
}