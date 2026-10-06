package main

import (
	"gin-learn/db"

	"gin-learn/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	server:=gin.Default()
	routes.RegisterRoutes(server)
	db.InitDB()
	server.Run(":8080")
}
