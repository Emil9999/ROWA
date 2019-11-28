package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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
	database, err := sql.Open("sqlite3", "/home/marcel/go/src/rowa/backend/mock.db")
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = database
	s.store = &Database{Db: database}
	err = s.store.DbSetup()
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	// Close the connection after all tests in the suite finish
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestGetLastSensorEntry() {
	// Insert two Entries in the Sensor DB
	_, err := s.db.Exec(`INSERT INTO SensorMeasurements (Datetime, Temp, LightIntensity)
 								VALUES('2018-06-05', 2.3, 120)`)
	_, err = s.db.Exec(`INSERT INTO SensorMeasurements (Datetime, Temp, LightIntensity)
 								VALUES('2018-06-06', 2.5, 110)`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get last entry from Store Method
	lastSensorEntry, err := s.store.GetLastSensorEntry()
	if err != nil {
		s.T().Fatal(err)
	}

	// Assert that the details of the SensorEntry is the same as the one we inserted
	expectedEntry := &SensorData{"2018-06-06", 2.5, 110}
	assert.Equal(s.T(), lastSensorEntry, expectedEntry)
}

func (s *StoreSuite) TestGetPlantsPerType() {
	harvestablePlants, err := s.store.GetPlantsPerType("harvestable")
	if err != nil {
		s.T().Fatal(err)
	}

	expected := []*PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}
	assert.Equal(s.T(), harvestablePlants, expected)

	plantablePlants, err := s.store.GetPlantsPerType("plantable")
	if err != nil {
		s.T().Fatal(err)
	}

	expected = []*PlantsPerPlantType{{"Basil", 0}, {"Lettuce", 0}}
	assert.Equal(s.T(), plantablePlants, expected)
}
