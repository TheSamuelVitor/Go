package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Definition of a type structure
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type member struct {
	Id_member   string `json:"id_member"`
	Name_member string `json:"name_member"`
	Age         string `json: "age"`
}

type task struct {
	Id_task   string `json: "id_task"`
	Name_task string `json: "name_task"`
	Time      string `json: "time"`
}

type project struct {
	Id_project   string `json: "id_project"`
	Name_project string `json: "name_project"`
	Tasks        string `json: tasks`
}

var members = []member{
	{Id_member: "1", Name_member: "Samuel Vitor", Age: "21"},
	{Id_member: "2", Name_member: "Dayanne", Age: "21"},
}

var tasks = []task{
	{Id_task: "1", Name_task: "Aprender GET", Time: "Thursday"},
	{Id_task: "2", Name_task: "Aprender POST", Time: "Friday"},
}

var projects = []project{
	{Id_project: "1", Name_project: "Fazer API", Tasks: "1 e 2"},
}

// ---------------------------
// 		FUNCOES DOS TIMES
// ---------------------------
func getmembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, members)
}

func postmembers(c *gin.Context) {
	var newMember member

	if err := c.BindJSON(&newMember); err != nil {
		return
	}

	members = append(members, newMember)
	c.IndentedJSON(http.StatusCreated, newMember)
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

// ---------------------------
// 		FUNCOES TASKS
// ---------------------------

func getTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskbyId(c*gin.Context){
	id:= c.Param("id")

	for _,a := range tasks{
		if a.Id_task == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "task not found",
	})
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

func getProjectbyId(c*gin.Context){
	id:= c.Param("id")

	for _,a := range projects{
		if a.Id_project == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "project not found",
	})
}

// ---------------------------
// 	 FUNCAO DA PAGINA INICIAL
// ---------------------------

var welcome = []string{
	"",
	"Bem vindo",
	"",
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

	server.GET("/get/members", getmembers)
	server.GET("/get/task", getTask)
	server.GET("/get/projects", getProjects)

	server.GET("/get/members/:id", getmembersbyId)
	server.GET("get/task/:id", getTaskbyId)
	server.GET("/get/projects/:id", getProjectbyId)

	server.POST("/post/members", postmembers)
	server.POST("/post/task", postTask)
	server.POST("/post/projects", postProjects)

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
