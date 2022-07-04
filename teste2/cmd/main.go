package main

import (
	"github.com/TheSamuelVitor/Go/tree/main/teste2/pkg/books"
	"github.com/TheSamuelVitor/Go/tree/main/teste2/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r,h)

	r.Run(port)
}