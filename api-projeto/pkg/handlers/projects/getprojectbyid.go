package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func GetProjectbyId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range mocks.Projects {
		if a.Id_project == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "project not found",
	})
}
