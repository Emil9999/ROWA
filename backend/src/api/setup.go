package api

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func DbSetup() {
	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")
	//Creating PlantType DB
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS PlantType (Name TEXT PRIMARY KEY, GrowthTime INTEGER)") // TODO Why no ID?
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO PlantType (Name, Growthtime) VALUES (?, ?)")

	statement.Exec("Lettuce", 42) //TODO Why only one PlantType, whether we use two plants?
	statement.Exec("Basil", 21)
	rows, _ := database.Query("SELECT Name, Growthtime from PlantType")
	var name string
	var growthTime int
	for rows.Next() {
		rows.Scan(&name, &growthTime)
		fmt.Println(name + ": " + strconv.Itoa(growthTime))
	}
	//Creating Module DB
	statement, _ = database.Prepare("DROP TABLE IF EXISTS Module")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS Module (Id INTEGER PRIMARY KEY, PlantType TEXT, Position INTEGER UNIQUE, AvailableSpots INTEGER, TotalSpots INTEGER, FOREIGN KEY(PlantType) REFERENCES PlantType(Name))")
	statement.Exec()
	statement, _ = database.Prepare("INSERT OR IGNORE INTO Module (Position, PlantType, AvailableSpots, TotalSpots) VALUES (?, ? ,?, ?)")
	statement.Exec(1, "Basil", 0, 6)
	statement.Exec(2, "Basil", 0, 6)
	statement.Exec(3, "Lettuce", 0, 6)
	statement.Exec(4, "Lettuce", 0, 6)
	statement.Exec(5, "Lettuce", 0, 6)
	statement.Exec(6, "Basil", 0, 6)

	// Creating Plant DB
	statement, _ = database.Prepare("DROP TABLE IF EXISTS Plant")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS Plant (Id INTEGER PRIMARY KEY, Module INTEGER, PlantPosition INTEGER, PlantDate TEXT, Harvested INTEGER, FOREIGN KEY (Module) REFERENCES Module(Id))")
	statement.Exec()
	statement, _ = database.Prepare("INSERT OR IGNORE INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)")

	// Add multiple Plants for testing purposes
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			statement.Exec(i+1, j+1, time.Now().Format("2006-01-02"), 0)
		}
	}

	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-09'  WHERE  Id = 6")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-08'  WHERE Id = 12")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-10'  WHERE Id = 24")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-30'  WHERE Id = 30")
	statement.Exec()

	//Make  Empty Plant Pot
	/*statement, _ = database.Prepare("UPDATE Plant SET Harvested = 1  WHERE Id = 24")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET Harvested = 1  WHERE Id = 11")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET Harvested = 1  WHERE Id = 12")
	statement.Exec()*/

	rows, _ = database.Query("SELECT Id, Position, PlantType, AvailableSpots, TotalSpots from Module")
	var Position int
	var Id int
	var PlantType string
	var AvailableSpots int
	var TotalSpots int
	for rows.Next() {
		rows.Scan(&Id, &Position, &PlantType, &AvailableSpots, &TotalSpots)
		fmt.Println(strconv.Itoa(Id), strconv.Itoa(Position)+PlantType+strconv.Itoa(AvailableSpots)+strconv.Itoa(TotalSpots))
	}
}
