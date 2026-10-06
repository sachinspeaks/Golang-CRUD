package routes

import (
	"gin-learn/models"
	"net/http"

	"gin-learn/utils"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context){
	var user models.USER
	err:=ctx.ShouldBindJSON(&user)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"Could not parse request data, check the data sent by you.",
		})
		return
	}
	err=user.Save()

	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"Error while saving user, please try again later." +err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":"User created successfully",
	})
}

func Login(ctx *gin.Context){
	var user models.USER
	err:=ctx.ShouldBindJSON(&user)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"Body not valid.",
		})
		return
	}
	err=user.ValidateCredentials()
	if err!=nil{
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"message":"User authentication failed."+err.Error(),
		})
		return
	}
	var token string
	token,err=utils.GenerateToken(user.Email,user.ID)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"User authentication failed."+err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK,gin.H{
		"message":"Login successful.",
		"token":token,
	})

}