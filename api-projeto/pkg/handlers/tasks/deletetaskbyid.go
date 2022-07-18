package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func DeleteTaskbyId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range mocks.Tasks {
		if a.Id_task == id {
			mocks.Tasks = append(mocks.Tasks[:i], mocks.Tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, mocks.Tasks)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
	}
}
