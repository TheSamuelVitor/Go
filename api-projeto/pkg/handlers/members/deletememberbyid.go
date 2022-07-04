package handlers

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func DeleteMembersbyId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range mocks.Members {
		if a.Id_member == id {
			mocks.Members = append(mocks.Members[:i], mocks.Members[i+1:]...)
			c.IndentedJSON(http.StatusOK, mocks.Members)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "member not found"})
	}
}