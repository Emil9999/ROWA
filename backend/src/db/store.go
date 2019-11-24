package db

import "database/sql"

type Store interface {
	GetPlantsPerType(p string) ([]*PlantsPerPlantType, error)
	GetLastSensorEntry() (*SensorData, error)
	GetHarvestablePlant(*PlantType) (*PositionOnFarm, error)
	HarvestDone(*PositionOnFarm) (*Status, error)
	Plant(*PlantType) (int, error)
	FinishPlanting(*PlantedModule) error
	DbSetup()
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type Database struct {
	Db *sql.DB
}

// The FunctionStore variable is a package level variable that will be available for
// use throughout our application code
var FunctionStore Store

/*
We will need to call the InitStore method to initialize the FunctionStore. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the FunctionStore as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	FunctionStore = s
}
