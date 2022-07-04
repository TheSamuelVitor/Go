package main

import (
	"github.com/TheSamuelVitor/Go/tree/main/teste2/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main()  {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	server := gin.Default()
	db.Init(dbUrl)
	server.GET("/", func (c* gin.Context){
		c.JSON(200, gin.H{
			"port" : port,
			"dbURL" : dbUrl,
		})
	})

	server.Run(port)
}