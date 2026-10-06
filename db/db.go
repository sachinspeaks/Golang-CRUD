package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
	var err error 
	DB,err=sql.Open("sqlite3","events.db")
	if err!=nil{
		log.Fatal(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	
	CreateTable()
	fmt.Println("Database Connected")
}

func CreateTable(){
	createUserTable:=`
	CREATE TABLE IF NOT EXISTS user (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_,err:=DB.Exec(createUserTable)
	if err!=nil{
		panic("Could not create the user table."+err.Error())
	}

	createEventsTable:=`
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER ,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_,err=DB.Exec(createEventsTable)
	if err!=nil{
		panic("Could not create the events table.")
	}

	createRegistrationTable:=`
	CREATE TABLE IF NOT EXISTS registration (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES user(id)
	)
	`
	_,err=DB.Exec(createRegistrationTable)
	if err!=nil{
		panic("Could not create the registrations table."+err.Error())
	}

}