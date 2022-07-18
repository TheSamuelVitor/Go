package members

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func Getmembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mocks.Members)
}
