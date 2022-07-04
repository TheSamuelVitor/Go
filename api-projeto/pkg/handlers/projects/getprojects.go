package handlers

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mocks.Projects)
}