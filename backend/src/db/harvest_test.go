package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func (s *StoreSuite) TestGetHarvestablePlant() {

	inputs := []*PlantType{{Name: "Basil"}, {Name: "Lettuce"}}
	expectedOutcome := []*PositionOnFarm{{6, 1}, {6, 4}}

	for i, element := range inputs {
		harvestablePlant, err := s.store.GetHarvestablePlant(element)
		if err != nil {
			s.T().Fatal(err)
		}

		assert.Equal(s.T(), harvestablePlant, expectedOutcome[i])
	}
}

func (s *StoreSuite) TestHarvestDone() {
	harvestDone, err := s.store.HarvestDone(&PositionOnFarm{6, 1})
	if err != nil {
		s.T().Fatal(err)
	}
	expected := &Status{"harvest done"}
	assert.Equal(s.T(), harvestDone, expected)

	plantablePlants, err := s.store.GetPlantsPerType("plantable")
	if err != nil {
		s.T().Fatal(err)
	}

	dbChange := []*PlantsPerPlantType{{"Basil", 1}, {"Lettuce", 0}}
	assert.Equal(s.T(), plantablePlants, dbChange)
}
