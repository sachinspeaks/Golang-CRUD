package routes

import (
	"gin-learn/middlewares"

	"github.com/gin-gonic/gin"
)
func RegisterRoutes(server *gin.Engine){
	server.GET("/events",GetEvents)

	authenticated:=server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events",CreateEvents)
	authenticated.PUT("/events/:id",UpdateEvent)
	authenticated.DELETE("/events/:id",DeleteEvent)
	authenticated.POST("/events/:id/register",RegisterForEvent)
	authenticated.DELETE("/events/:id/register",CancelRegistration)

	server.GET("/events/:id",GetEvent)
	server.POST("/signup",Signup)
	server.POST("/login",Login)
}