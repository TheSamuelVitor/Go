package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// Definition of a type structure

type member struct {
	Id_member   string `json:"id_member"`
	Name_member string `json:"name_member"`
	Age         string `json: "age"`
	Team        string `json: "name_team"`
}

type team struct {
	Id_team   string `json: "id_task"`
	Name_team string `json: "name_team"`
	task      string `json: "id_task"`
}

type task struct {
	Id_task   string `json: "id_task"`
	Name_task string `json: "name_task"`
	Time      string `json: "id_time"`
	Id_member string `json: "id_member"`
}

type project struct {
	Id_project   string `json: "id_project"`
	Name_project string `json: "name_project"`
	Tasks        string `json: "id_task"`
}

var members = []member{
	{Id_member: "1", Name_member: "Samuel Vitor", Age: "21", Team: "Komanda"},
	{Id_member: "2", Name_member: "Dayanne", Age: "21", Team: "Komanda"},
}

var teams = []team{
	{Id_team: "1", Name_team: "Komanda", task: "Aprender GET"},
	{Id_team: "2", Name_team: "DevsKariri", task: "Aprender POST"},
}

var tasks = []task{
	{Id_task: "1", Name_task: "Aprender GET", Time: "Thursday", Id_member: "1"},
	{Id_task: "2", Name_task: "Aprender POST", Time: "Friday", Id_member: "1"},
}

var projects = []project{
	{Id_project: "1", Name_project: "Fazer API", Tasks: "1 e 2"},
}

// ---------------------------
// 	  FUNCOES DOS MEMBROS
// ---------------------------

func getmembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, members)
}

func getmembersbyId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range members {
		if a.Id_member == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "member not found"})
}

func postmembers(c *gin.Context) {
	var newMember member

	if err := c.BindJSON(&newMember); err != nil {
		return
	}

	members = append(members, newMember)
	c.IndentedJSON(http.StatusCreated, newMember)
}

func putMembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "putMembers",
	})
}

func deleteMembersbyId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range members {
		if a.Id_member == id {
			members = append(members[:i], members[i+1:]...)
			c.IndentedJSON(http.StatusOK, members)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "member not found"})
	}
}

// ---------------------------
// 		FUNCOES TEAMS
// ---------------------------

func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
}

func getTeambyId(c *gin.Context) {
	id := c.Param("id")
	for _, a := range teams {
		if a.Id_team == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "team not found",
	})
}

// ---------------------------
// 		FUNCOES TASKS
// ---------------------------

func getTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskbyId(c *gin.Context) {
	id := c.Param("id")
	for _, a := range tasks {
		if a.Id_task == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "task not found",
	})
}

func postTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func putTask(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK, gin.H{
			"message": "put task",
		})
}

func deleteTaskbyId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range tasks {
		if a.Id_task == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, tasks)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
	}
}

// ---------------------------
// 		FUNCAO DOS PROJETOS
// ---------------------------
func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func postProjects(c *gin.Context) {
	var newProject project

	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}

func putProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "put Projects",
	})
}

func getProjectbyId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range projects {
		if a.Id_project == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "project not found",
	})
}

func deleteProjectbyId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range projects {
		if a.Id_project == id {
			projects = append(projects[:i], projects[i+1:]...)
			c.IndentedJSON(http.StatusOK, projects)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
	}
}

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

// func getTaskbyMember()  {

// }

// ---------------------------
// 		FUNCAO PRINCIPAL
// ---------------------------

func main() {
	server := gin.Default()

	server.GET("/", bemvindo)

	server.GET("/members", getmembers)
	server.GET("/teams", getTeams)
	server.GET("/tasks", getTask)
	server.GET("/projects", getProjects)

	server.GET("/members/:id", getmembersbyId)
	server.GET("/teams/:id", getTeambyId)
	server.GET("/tasks/:id", getTaskbyId)
	server.GET("/projects/:id", getProjectbyId)

	server.POST("/members", postmembers)
	server.POST("/tasks", postTask)
	server.POST("/projects", postProjects)

	server.PUT("/members", putMembers)
	server.PUT("/tasks", putTask)
	server.PUT("/projects", putProjects)

	server.DELETE("/members/:id", deleteMembersbyId)
	server.DELETE("/tasks/:id", deleteTaskbyId)
	server.DELETE("/projects/:id", deleteProjectbyId)

	/*

		GET    = used to retrieve data from the server, read-only method.
		POST   = used to send data to the server, associate with an ID, used to create new data entry.
		PUT    = used to update a existing resource.
		DELETE = used to delete a resource specified in the URL.

		server.GET("/someGet", getting)
		server.POST("/somePost", posting)
		server.PUT("/somePut", putting)
		server.DELETE("/someDelete", deleting)

		port := os.Getenv("PORT")
		server.Run(":"+port)

	*/

	server.Run(":8080")
}
