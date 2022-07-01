package main

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
	_"github.com/lib/pq"
)

var db *sql.DB
var err error

type member struct {
	Id_member   string `json:"id_member"`
	Name_member string `json:"name_member"`
	Id_Team     string `json:"id_team"`
}

type team struct {
	Id_team   string `json:"id_team"`
	Name_team string `json:"name_team"`
}

type task struct {
	Id_task    string `json:"id_task"`
	Name_task  string `json:"name_task"`
	Id_member  string `json:"id_member"`
	Id_project string `json:"id_project"`
}

type project struct {
	Id_project   string `json:"id_project"`
	Name_project string `json:"name_project"`
}

var members = []member{
	{Id_member: "1", Name_member: "Samuel Vitor", Id_Team: "1"},
	{Id_member: "2", Name_member: "Dayanne", Id_Team: "1"},
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
// 	  GENERAL GETS FUNCTIONS
// ---------------------------

func getmembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, members)
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
}

func getTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// ---------------------------
// 	 SPECIFICS GET FUNCTIONS
// ---------------------------

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

// -------------------------------
// 		GET FUNCTIONS BY ID
// -------------------------------
func getMembersbyTeamid(c *gin.Context) {
	id := c.Param("id")

	for _, a := range members {
		if a.Id_Team == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "none member found",
	})
}

func getTasksbyProjectid(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tasks {
		if a.Id_project == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "none task found",
	})
}

// ---------------------------
// 	 	POST FUNCTIONS
// ---------------------------

func postmembers(c *gin.Context) {
	var newMember member

	if err := c.BindJSON(&newMember); err != nil {
		return
	}

	members = append(members, newMember)
	c.IndentedJSON(http.StatusCreated, newMember)
}

func postTeam(c *gin.Context) {
	var newTeam team

	if err := c.BindJSON(&newTeam); err != nil {
		return
	}

	teams = append(teams, newTeam)
	c.IndentedJSON(http.StatusCreated, newTeam)
}

func postTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func postProjects(c *gin.Context) {
	var newProject project

	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}

// ---------------------------
// 		PUT FUNCTIONS
// ---------------------------

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

// ---------------------------
// 		FUNCAO DELETE
// ---------------------------

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
// 	 	MAIN FUNCTION
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

	server.GET("/members/team/:id", getMembersbyTeamid)
	server.GET("/tasks/project/:id", getTasksbyProjectid)

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

	server.Run(":3000")
}
