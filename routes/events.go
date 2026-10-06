package routes

import (
	"gin-learn/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetEvent(ctx *gin.Context){
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
			"message":"Something went wrong, please try again later."+err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,event)
}

func GetEvents(ctx *gin.Context){
	events,err:=models.GetAllEvents()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not fetch events, try again later.",
		})
	}
	ctx.JSON(http.StatusOK,events)
}

func CreateEvents(ctx *gin.Context){
	var event models.Event
	err:=ctx.ShouldBindJSON(&event)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"Some required field was not provided, check the data sent by you.",
		})
		return
	}
	event.UserID=ctx.GetInt64("userId")
	err=event.Save()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not create event, try again later."+err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message":"Event created successfully.",
		"event":event,
	})
}

func UpdateEvent(ctx *gin.Context){
	eventId,err:=strconv.ParseInt(ctx.Param("id"),10,64)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"Could not parse event Id, please try again later.",
		})
		return 
	} 
	var event models.Event
	event,err=models.GetEventById(eventId)
	if event.UserID!=ctx.GetInt64("userId"){
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"message":"You are not authorized to udpate this event",
		})
		return
	}
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not fetch the event, please try again later.",
		})
		return 
	}
	var UpdatedEvent models.Event
	err=ctx.ShouldBindJSON(& UpdatedEvent)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"You've sent a erroneous body.",
		})
		return
	}
	UpdatedEvent.Id=eventId
	err=UpdatedEvent.Update()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not update event, please try again later.",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Event updated successfully.",
	})
}

func DeleteEvent(ctx *gin.Context){
	eventId,err:=strconv.ParseInt(ctx.Param("id"),10,64)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"Could not parse event Id, please try again later.",
		})
		return 
	}

	event,err2:=models.GetEventById(eventId)
	if err2!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not fetch the event, please try again later.",
		})
		return 
	}
	if event.UserID!=ctx.GetInt64("userId"){
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"message":"You are not authorized to delete this event",
		})
		return
	}
	err=event.Delete()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not delete the event, please try again later.",
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Event deleted successfully.",
	})
}