package members

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func GetmembersbyId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range mocks.Members {
		if a.Id_member == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "member not found"})
}