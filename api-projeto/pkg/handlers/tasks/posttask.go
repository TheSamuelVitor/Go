package handlers

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/mocks"
	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/models"
	"github.com/gin-gonic/gin"
)

func PostTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	mocks.Tasks = append(mocks.Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}