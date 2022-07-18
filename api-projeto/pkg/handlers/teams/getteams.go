package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/mocks"
	"github.com/gin-gonic/gin"
)


func GetTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mocks.Teams)
}
