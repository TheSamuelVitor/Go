package handlers

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func GetTaskbyId(c *gin.Context) {
	id := c.Param("id")
	for _, a := range mocks.Tasks {
		if a.Id_task == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "task not found",
	})
}