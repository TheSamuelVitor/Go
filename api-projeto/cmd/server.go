package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// ---------------------------
// 	 FUNCAO DA PAGINA INICIAL
// ---------------------------

var welcome = []string{
	"",
	"Bem vindo",
	"",
	"--- ROTAS --- ",
	"GET    /",
	"GET    /members",
	"GET    /task",
	"GET    /projects",
	"GET    /members/:id",
	"GET    /task/:id",
	"GET    /projects/:id",
	"",
	"POST   /members",
	"POST   /task",
	"POST   /projects",
	"",
	"DELETE /members/:id",
	"DELETE /tasks/:id",
	"DELETE /project/:id",
}

func bemvindo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, welcome)
}

// ---------------------------
// 		FUNCAO PRINCIPAL
// ---------------------------

func main() {
	server := gin.Default()

	server.GET("/", bemvindo)

	server.GET("/members", handlers.Getmembers)
	server.GET("/members/:id", handlers.GetmembersbyId)
	server.POST("/members", handlers.Postmembers)
	server.PUT("/members", handlers.PutMembers)
	server.DELETE("/members/:id", handlers.DeleteMembersbyId)

	server.GET("/teams", handlers.GetTeams)
	server.GET("/teams/:id", handlers.GetTeambyId)

	server.GET("/tasks", handlers.GetTask)
	server.GET("/tasks/:id", handlers.GetTaskbyId)
	server.POST("/tasks", handlers.PostTask)
	server.PUT("/tasks", handlers.PutTask)
	server.DELETE("/tasks/:id", handlers.DeleteTaskbyId)

	server.GET("/projects", handlers.GetProjects)
	server.GET("/projects/:id", handlers.GetProjectbyId)
	server.POST("/projects", handlers.PostProjects)
	server.PUT("/projects", handlers.PutProjects)
	server.DELETE("/projects/:id", handlers.DeleteProjectbyId)

	server.Run(":8080")
}
