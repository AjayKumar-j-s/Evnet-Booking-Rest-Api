package models

import (
	"github.com/AjayKumar-j-s/EventPlanner/db"
)

type Reg struct {
	ID      int64
	Userid  int64
	Eventid int64
}

func (r *Reg) Register()(error) {

	query := `INSERT INTO Registration(UserID,EventID)
	values(?,?)`

	_,err := db.DB.Exec(query,r.Userid,r.Eventid)

	return err

}