package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func DeleteProjectbyId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range mocks.Projects {
		if a.Id_project == id {
			mocks.Projects = append(mocks.Projects[:i], mocks.Projects[i+1:]...)
			c.IndentedJSON(http.StatusOK, mocks.Projects)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
	}
}