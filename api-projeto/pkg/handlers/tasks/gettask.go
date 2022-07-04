package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mocks.Tasks)
}
