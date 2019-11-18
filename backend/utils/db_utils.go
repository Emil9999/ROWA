package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DbObject struct {
	Database *sql.DB
}

func (db *DbObject) Init() {
	statement, _ := db.Database.Prepare("CREATE TABLE IF NOT EXISTS PlantType (Name TEXT PRIMARY KEY, GrowthTime INTEGER)")
	statement.Exec()
}

func GetHarvestablePlantQuery(plantType string, db *DbObject) []int {
	sqlQuery := fmt.Sprintf("SELECT PlantPosition, Module FROM Plant INNER JOIN Module M on Plant.Module = M.Id INNER JOIN PlantType PT on M.PlantType = PT.Name where Harvested = 0 and date(PlantDate, '+'||GrowthTime||' days') <= date('now') AND PlantType = '%s'", plantType)
	rows, err := db.Database.Query(sqlQuery)
	//We only need the position (?)
	var position int
	var module int
	var arr []int
	if err != nil {
		fmt.Println("db query err: ", err)
	}
	for rows.Next() {
		rows.Scan(&position, &module)
		arr = append(arr, position)
		arr = append(arr, module)
	}
	arr = arr[:2]
	return arr

}

func UpdateAfterHarvest(plantPosition int, modulePosition int, db *DbObject) bool {

	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")
	sqlQuery := fmt.Sprintf("UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = '%d' AND Module=%d", plantPosition, modulePosition)
	statement, _ := database.Prepare(sqlQuery)
	statement.Exec()
	sqlQuery = fmt.Sprintf("UPDATE Module SET AvailableSpots = AvailableSpots + 1 WHERE Position='%d'", modulePosition)
	statement, _ = database.Prepare(sqlQuery)
	statement.Exec()
	return true
}
