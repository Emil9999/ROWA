package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DbObject struct {
	database *sql.DB
}

func (db *DbObject) Init() {
	statement, _ := db.database.Prepare("CREATE TABLE IF NOT EXISTS PlantType (Name TEXT PRIMARY KEY, GrowthTime INTEGER)")
	statement.Exec()
}

func main() {
	database, err := sql.Open("sqlite3", "./rowa.db")
	if err != nil {
		fmt.Println(err)
	}
	db := &DbObject{database}
	db.Init()
}
