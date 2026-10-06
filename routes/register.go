package routes

import (
	"gin-learn/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context){
	userId:=ctx.GetInt64("userId")
	eventId,err:=strconv.ParseInt(ctx.Param("id"),10,64)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"Could not parse event Id, please try again later.",
		})
		return 
	} 

	var event models.Event
	event,err=models.GetEventById(eventId)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not fetch event.",
		})
		return
	}

	err=event.Register(userId)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
		"message":"Could not register user for event.",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":"User registered successfully.",
	})
}

func CancelRegistration(ctx *gin.Context){
	userId:=ctx.GetInt64("userId")
	eventId,err:=strconv.ParseInt(ctx.Param("id"),10,64)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"Could not parse event Id, please try again later.",
		})
		return 
	} 
	var event models.Event
	event.Id=eventId

	err=event.CancelRegistration(userId)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not cancel registration, please try again later.",
		})
		return 
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Registration cancelled successfully.",
	})

}