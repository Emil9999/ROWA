package db

import (
	"github.com/stretchr/testify/assert"
)

func (s *StoreSuite) TestPlant() {
	s.store.CreatePlantTypes()
	s.store.CreateModuleWithFreeSpot(1, "Basil", 2)

	_, err := s.store.HarvestDone(&PositionOnFarm{6, 1})
	if err != nil {
		s.T().Fatal(err)
	}

	modulePosition, err := s.store.Plant(&PlantType{Name: "Basil"})
	if err != nil {
		s.T().Fatal(err)
	}

	expected := 1
	assert.Equal(s.T(), modulePosition, expected)
}

func (s *StoreSuite) TestFinishPlanting() {
	s.store.CreatePlantTypes()
	s.store.CreateModuleWithFreeSpot(1, "Basil", 2)

	_, err := s.store.HarvestDone(&PositionOnFarm{6, 1})
	if err != nil {
		s.T().Fatal(err)
	}

	message, _ := s.store.FinishPlanting(&PlantedModule{Module: 1})

	expected := &Status{Message: "Planting Done"}
	assert.Equal(s.T(), message, expected)
}
