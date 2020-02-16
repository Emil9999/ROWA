package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	
)

type Times struct{
	TimeOn  int `json:"time_on"`
	TimeOff int `json:"time_off"`
}

type CurrentTimes struct{
	TimeOn  int `json:"time_on"`
	TimeOff int `json:"time_off"`
	State int `json:"state"`
}


//store incoming new time data 
func (store* Database) InsertLightTimes(newTimes *Times) (status *Status, err error){

	//Calculate Values from icomming time 

		//Insert into db

		//LightTimerStart()

return
}

//get time and state from db
func (store* Database) GetLightTimes() (currenTimes *CurrentTimes, err error){

	//get Times and State from DB
	//make Json and send

return
}


//get times at server start and return
func (store* Database) LightTimerStart(){
	sqlQuery := `SELECT OnTime, OffTime
				 FROM TimeTable
				 WHERE ID = 1`

		//Start go routines with values

	return
}





// wholeNumber(x/60) = h x- (wholeNumber(x/60))*60 = min