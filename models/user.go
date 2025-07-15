package models

import (
	"errors"

	"github.com/AjayKumar-j-s/EventPlanner/db"
	"github.com/AjayKumar-j-s/EventPlanner/utils"
)

type User struct {
	UserID   int64
	Email    string
	Password string
}

func (u User) Save() error {

	query := `INSERT INTO User (Email,Password) values (?,?)`

	stmt, err := db.DB.Prepare(query)

	if(err != nil){
		return err
	}

	defer stmt.Close()

	Hpass,err := utils.HashPassword(u.Password)

	if(err != nil){
		return err
	}

	_,err = stmt.Exec(u.Email,Hpass)

	return err



}

func (u User) ValidatePassword()(error){

	query := `SELECT Password FROM User WHERE Email = ?`

	stmt,err := db.DB.Prepare(query)

	if(err != nil){
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRow(u.Email)

	var Haspass string

	err = row.Scan(&Haspass)
	if(err !=  nil){
		return err
	}

	PasswordIsValid := utils.CheckHash(u.Password,Haspass)

	if(!PasswordIsValid){
		return errors.New("invalid Credentials")
	}

	return nil


}