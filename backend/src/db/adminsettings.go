package db

import (

	_ "github.com/mattn/go-sqlite3"
	"util"
	
)

type Times struct{
	TimeOn  string `json:"time_on"`
	TimeOff string `json:"time_off"`
}

type CurrentTime struct{
	TimeOn  string `json:"time_on"`
	TimeOff string `json:"time_off"`
	State   int `json:"state"`
}

type LightState struct{
	State int `json:"state"`
}


//store incoming new time data 
func (store* Database) InsertLightTimes(newTimes *Times) (status *Status, err error){

	//TODO Calculate Values from icomming time 
	status = &Status{}

	sqlQuery := `UPDATE TimeTable SET OnTime = ?, OffTime = ? WHERE ID = 1`
	statement, _ := store.Db.Prepare(sqlQuery)
	defer statement.Close()

	_, err = statement.Exec(newTimes.TimeOn, newTimes.TimeOff)
	if err != nil {
		status.Message = "error"
		return
	}
	 util.LightTimesRenew()


	status.Message = "harvest done"
	return
}

//get time and state from db
func (store* Database) GetLightTimes() (currentTime *CurrentTime, err error){
	
	sqlQuery := `SELECT OnTime, OffTime, CurrentState
				 FROM TimeTable
				 WHERE ID = 1`



	rows, err := store.Db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return
	}

	currentTime = &CurrentTime{}
	rows.Next()
	err = rows.Scan(&currentTime.TimeOn, &currentTime.TimeOff, &currentTime.State )

	if err != nil {
		return
	}
	//get Times and State from DB
	//make Json and send

return
}




// wholeNumber(x/60) = h x- (wholeNumber(x/60))*60 = min