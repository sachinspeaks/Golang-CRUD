package models

import (
	"errors"
	"fmt"
	"gin-learn/db"
	"gin-learn/utils"
)

type USER struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u USER)Save()error{
	query:=`INSERT INTO user(email,password) VALUES (?,?)`
	statement,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer statement.Close()
	hashedPassword,err:=utils.HashPassword(u.Password)
	if err!=nil{
		return err
	}
	result,err2:=statement.Exec(u.Email,hashedPassword)
	if err2!=nil{
		return err2
	}
	userId,err3:=result.LastInsertId()
	u.ID=userId
	return err3
}

func (u *USER)ValidateCredentials() error{
	query:=`SELECT id,password FROM user WHERE email=?`
	row:=db.DB.QueryRow(query,u.Email)
	var retrievedPassword string
	
	err:=row.Scan(&u.ID,&retrievedPassword)
	if err!=nil{
		fmt.Println(err.Error())
		return errors.New("Invalid Credentials.")
	}
	isPasswordValid:=utils.CheckPasswordHash(u.Password,retrievedPassword)
	if !isPasswordValid{
		return errors.New("Invalid Credentials.")
	}
	return nil
}