package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func GetTeambyId(c *gin.Context) {
	id := c.Param("id")
	for _, a := range mocks.Teams {
		if a.Id_team == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "team not found",
	})
}