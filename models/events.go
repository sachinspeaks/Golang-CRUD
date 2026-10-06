package models

import (
	"gin-learn/db"
	"time"
)

type Event struct{
	Id int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time 
	UserID int64
}

var events []Event=[]Event{}

func (e *Event) Save() error {
	query:=`
	INSERT INTO events(name,description,location,dateTime,user_id)
	values (?,?,?,?,?)`
	statement,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer statement.Close()
	result,err2:=statement.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	if err2!=nil{
		return err2
	}
	var id int64
	id,err=result.LastInsertId()
	e.Id=id+1
	return err
}

func GetAllEvents() ([]Event,error){
	query:=`
	select * from events
	`
	rows,err:=db.DB.Query(query)
	defer rows.Close()
	if err!=nil{
		return nil,err
	}
	var events []Event
	for rows.Next(){
		var event Event
		err:=rows.Scan(&event.Id,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
		if err!=nil{
			return nil,err
		}
		events=append(events, event)
	}
	return events,nil
}

func GetEventById(id int64) (Event,error){
	query:=`
	SELECT * FROM events WHERE ID=?
	`
	row:=db.DB.QueryRow(query,id)
	var event Event
	err:=row.Scan(&event.Id,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
	if err!=nil{
		return Event{},err
	}
	return event,nil
}

func (e Event)Update() error{
	query:=`
	UPDATE events SET name=?,description=?,location=?,dateTime=? WHERE id=?
	`
	statement,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer statement.Close()
	_,err=statement.Exec(e.Name,e.Description,e.Location,e.DateTime,e.Id)
	return err
}

func (event Event) Delete() error{ 
	query:=`DELETE FROM events WHERE ID=?`
	statement,err:=db.DB.Prepare(query)

	if err!=nil{
		return err
	}
	defer statement.Close()
	_,err=statement.Exec(event.Id)
	return err
}

func (e Event) Register(userId int64)error{
	query:=`INSERT INTO registration(event_id,user_id) values (?,?)`
	statement,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer statement.Close()
	_,err=statement.Exec(e.Id,userId)
	return err
}

func (e Event)CancelRegistration(userId int64)error{
	query:=`
	DELETE FROM registrations WHERE event_id=? and user_id=?
	`
	statement,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer statement.Close()
	_,err=statement.Exec(e.Id,userId)
	return err
}