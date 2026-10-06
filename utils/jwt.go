package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret="q32aefgq3245effq345gaef#$%agffq#$%"

func GenerateToken(email string,userId int64) (string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})
	signedString,err:= token.SignedString([]byte(secret))
	if err!=nil{
		return "",err
	}
	return signedString,nil
}

func VerifyToken(token string)(int64,error){
	parsedToken,err:=jwt.Parse(token,func (token *jwt.Token)(interface{},error){
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil,errors.New("Token matching algorithm problem.")
		}
		return []byte(secret),nil
	})
	if err!=nil{
		return 0,errors.New("Could not parse token.")
	}
	isTokenValid:=parsedToken.Valid
	if !isTokenValid{
		return 0,errors.New("Token not valid.")
	}
	claims,ok:=parsedToken.Claims.(jwt.MapClaims)
	if !ok{
		return 0,errors.New("Invalid token claims.")
	}
	// email:=claims["email"].(string)
	userId:=int64(claims["userId"].(float64))
	return userId,nil

}