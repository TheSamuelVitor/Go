package main

import (
	"net/http"

	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/handlers/members"
	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/handlers/projects"
	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/handlers/tasks"
	"github.com/TheSamuelVitor/Go/tree/api-projeto/pkg/handlers/teams"
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

	server.GET("/members", members.Getmembers)
	server.GET("/members/:id", members.GetmembersbyId)
	server.POST("/members", members.Postmembers)
	server.PUT("/members", members.PutMembers)
	server.DELETE("/members/:id", members.DeleteMembersbyId)

	server.GET("/teams", teams.GetTeams)
	server.GET("/teams/:id", teams.GetTeambyId)

	server.GET("/tasks", tasks.GetTask)
	server.GET("/tasks/:id", tasks.GetTaskbyId)
	server.POST("/tasks", tasks.PostTask)
	server.PUT("/tasks", tasks.PutTask)
	server.DELETE("/tasks/:id", tasks.DeleteTaskbyId)

	server.GET("/projects", projects.GetProjects)
	server.GET("/projects/:id", projects.GetProjectbyId)
	server.POST("/projects", projects.PostProjects)
	server.PUT("/projects", projects.PutProjects)
	server.DELETE("/projects/:id", projects.DeleteProjectbyId)

	server.Run(":8080")
}
