package members

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func Getmembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mocks.Members)
}
