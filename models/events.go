package models

import (
	"time"
	"github.com/AjayKumar-j-s/EventPlanner/db"
)


type Event struct{
	//if the post request doesn't contains any value we can through error by struct tags
	//else the value set to nil
	ID int64 
	Name string `binding:"required"`
	Description string
	Location string
	DateTime time.Time
	UserID int64
}


func (e *Event) Save()(error){

	query := `INSERT INTO events(name,description,location,datetime,userid)
	VALUES(?,?,?,?,?)`
	//? is a special syntax to prevent sql iinjection

	stmt,err := db.DB.Prepare(query)
	//we can directly use Query or Exec function but prepare help to perfom efficient in
	//certain cases and store the query in sql package
	
	if(err != nil){
		return err
	}
	
	
	res,err := stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	
	if(err != nil){
		return err
	}
	
	id,err := res.LastInsertId() 
	
	e.ID = id
	
	defer stmt.Close()

	return err
}

func GetEvents() ([]Event,error){
	query := `SELECT * FROM events`

	var events []Event

	rows,err := db.DB.Query(query)

	
	if(err != nil){
		return nil,err
	}
	
	defer rows.Close()
	
	for rows.Next(){
		var e Event

		err := rows.Scan(&e.ID,&e.Name,&e.Description,&e.Location,&e.DateTime,&e.UserID)

		if err != nil{

			return nil,err
		}
		
		events = append(events, e)

	}

	return events,nil

}

func GetEvent(id int64)(*Event,error){
	query := `SELECT * FROM events WHERE ID = ?`

	row := db.DB.QueryRow(query,id)

	var e Event

	err := row.Scan(&e.ID,&e.Name,&e.Description,&e.Location,&e.DateTime,&e.UserID)

	if(err != nil){
		return nil,err
	}

	return &e,nil
}

func (e Event) UpdateEve()(error){
	query := `UPDATE events
	SET Name = ?,Description = ?,Location = ?,DateTime = ?,UserID = ?
	WHERE ID = ?		
	`

	stmt,err := db.DB.Prepare(query)

	
	if(err != nil){
		return err
	}
	
	defer stmt.Close()
	_,err = stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID,e.ID)

	return err


}


func DeleteEvent(id int64)(error){

	query := `DELETE FROM events
		WHERE ID = ?
		`

	stmt ,err := db.DB.Prepare(query)
	if(err != nil){
		return err
	}

	defer stmt.Close()

	_,err = stmt.Exec(id)

	return err

}