package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/TheSamuelVitor/Go/pkg/models"
	"github.com/gin-gonic/gin"
)

func PostProjects(c *gin.Context) {
	var newProject models.Project

	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	mocks.Projects = append(mocks.Projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}
