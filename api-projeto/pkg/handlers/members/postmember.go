package handlers

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/TheSamuelVitor/Go/pkg/models"
	"github.com/gin-gonic/gin"
)

func Postmembers(c *gin.Context) {
	var newMember models.Member

	if err := c.BindJSON(&newMember); err != nil {
		return
	}

	mocks.Members = append(mocks.Members, newMember)
	c.IndentedJSON(http.StatusCreated, newMember)
}