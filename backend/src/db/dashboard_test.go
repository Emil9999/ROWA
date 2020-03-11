package db

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	/*
		The suite is defined as a struct, with the store and db as its
		attributes. Any variables that are to be shared between tests in a
		suite should be stored as attributes of the suite instance
	*/
	store *Database
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	/*
		The database connection is opened in the setup, and
		stored as an instance variable,
		as is the higher level `store`, that wraps the `db`
	*/
	database, err := sql.Open("sqlite3", "../../mock.db")
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = database
	s.store = &Database{Db: database}
}

func (s *StoreSuite) SetupTest() {
	tableNames := [4]string{"Module", "Plant", "PlantType", "SensorMeasurements"}

	for _, tableName := range tableNames {
		_, err := s.db.Exec("DELETE FROM " + tableName)
		if err != nil {
			s.T().Fatal(err)
		}
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestGetLastSensorEntry() {
	// Insert two Entries in the Sensor DB
	_, err := s.db.Exec(`INSERT INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH)
 								VALUES('2018-06-05', 2.3, 120, 40, 30, 18, 6)`)
	_, err = s.db.Exec(`INSERT INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH)
 								VALUES('2018-06-06', 2.5, 110, 45, 20, 16, 8)`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get last entry from Store Method
	lastSensorEntry, err := s.store.GetLastSensorEntry()
	if err != nil {
		s.T().Fatal(err)
	}

	// Assert that the details of the SensorEntry is the same as the one we inserted
	expectedEntry := &SensorData{"2018-06-06", 2.5, 110, 45, 20, 16, 8}
	assert.Equal(s.T(), lastSensorEntry, expectedEntry)
}

func (s *StoreSuite) TestGetPlantsPerTypeHarvest() {
	s.store.CreatePlantTypes()
	s.store.CreateModuleWithPlants(1, "Basil")
	s.store.CreateModuleWithPlants(2, "Lettuce")

	harvestablePlants, err := s.store.GetPlantsPerType("harvestable")
	if err != nil {
		s.T().Fatal(err)
	}

	expected := []*PlantsPerPlantType{{"Basil", 4}, {"Lettuce", 1}}
	assert.Equal(s.T(), harvestablePlants, expected)
}

func (s *StoreSuite) TestGetPlantsPerTypePlantable() {
	s.store.CreateModule(1, "Basil", 2)
	s.store.CreateModule(2, "Lettuce", 4)
	plantablePlants, err := s.store.GetPlantsPerType("plantable")
	if err != nil {
		s.T().Fatal(err)
	}

	expected := []*PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 4}}
	assert.Equal(s.T(), plantablePlants, expected)
}
