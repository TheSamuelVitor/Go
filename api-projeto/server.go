package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type member struct {
	Id_member   string `json:"id_member"`
	Name_member string `json:"name_member"`
	Team        string `json: "name_team"`
}

type team struct {
	Id_team   string `json: "id_task"`
	Name_team string `json: "name_team"`
}

type task struct {
	Id_task    string `json: "id_task"`
	Name_task  string `json: "name_task"`
	Id_member  string `json: "id_member"`
	Id_project string `json: id_project`
}

type project struct {
	Id_project   string `json: "id_project"`
	Name_project string `json: "name_project"`
}

var members = []member{
	{Id_member: "1", Name_member: "Samuel Vitor", Team: "Komanda"},
	{Id_member: "2", Name_member: "Dayanne", Team: "Komanda"},
}

var teams = []team{
	{Id_team: "1", Name_team: "Komanda"},
	{Id_team: "2", Name_team: "DevsKariri"},
}

var tasks = []task{
	{Id_task: "1", Name_task: "Aprender GET", Id_member: "1", Id_project: "1"},
	{Id_task: "2", Name_task: "Aprender POST", Id_member: "1", Id_project: "1"},
}

var projects = []project{
	{Id_project: "1", Name_project: "Fazer API"},
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
	id := c.Param("id")
	for i := range members {
		if members[i].Id_member == id {
			c.BindJSON(&members[i])
			c.IndentedJSON(http.StatusOK, members[i])
			return
		}
	}
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

func postTeam(c *gin.Context) {
	var newTeam team

	if err := c.BindJSON(&newTeam); err != nil {
		return
	}

	teams = append(teams, newTeam)
	c.IndentedJSON(http.StatusCreated, newTeam)
}

func putTeam(c *gin.Context) {
	id := c.Param("id")
	for i := range teams {
		if teams[i].Id_team == id {
			c.BindJSON(&teams[i])
			c.IndentedJSON(http.StatusOK, teams[i])
			return
		}
	}
}

func deleteTeambyId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range teams {
		if a.Id_team == id {
			teams = append(teams[:i], teams[i+1:]...)
			c.IndentedJSON(http.StatusOK, teams)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "team not found"})
	}
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
	id := c.Param("id")
	for i := range tasks {
		if tasks[i].Id_task == id {
			c.BindJSON(&tasks[i])
			c.IndentedJSON(http.StatusOK, tasks[i])
			return
		}
	}
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
	id := c.Param("id")
	for i := range projects {
		if projects[i].Id_project == id {
			c.BindJSON(&projects[i])
			c.IndentedJSON(http.StatusOK, projects[i])
			return
		}
	}
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
	" ------------------- ",
	"   POSSIBLE ROUTES   ",
	" ------------------- ",
	"",
	" -- GET --",
	" GET - /",
	" GET - /members",
	" GET - /teams",
	" GET - /tasks",
	" GET - /projects",
	" GET - /members/:id",
	" GET - /teams/:id",
	" GET - /tasks/:id",
	" GET - /projects/:id",
	"",
	" -- POST --",
	" POST - /members",
	" POST - /tasks",
	" POST - /teams",
	" POST - /projects",
	"",
	" -- PUT --",
	" PUT - /members/:id",
	" PUT - /teams/:id",
	" PUT - /tasks/:id",
	" PUT - /projects/:id",
	"",
	" -- DELETE --",
	" DELETE - /members/:id",
	" DELETE - /tasks/:id",
	" DELETE - /teams/:id",
	" DELETE - /projects/:id",
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
	server.POST("/teams", postTeam)
	server.POST("/projects", postProjects)

	server.PUT("/members/:id", putMembers)
	server.PUT("/teams/:id", putTeam)
	server.PUT("/tasks/:id", putTask)
	server.PUT("/projects/:id", putProjects)

	server.DELETE("/members/:id", deleteMembersbyId)
	server.DELETE("/tasks/:id", deleteTaskbyId)
	server.DELETE("teams/:id", deleteTeambyId)
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
