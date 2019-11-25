package db

import (
	"github.com/stretchr/testify/assert"
)

func (s *StoreSuite) TestPlant() {
	_, err := s.store.HarvestDone(&PositionOnFarm{6, 2})
	if err != nil {
		s.T().Fatal(err)
	}

	modulePosition, err := s.store.Plant(&PlantType{Name: "Basil"})
	if err != nil {
		s.T().Fatal(err)
	}

	expected := 2
	assert.Equal(s.T(), modulePosition, expected)
}

func (s *StoreSuite) TestFinishPlanting() {
	_, err := s.store.HarvestDone(&PositionOnFarm{6, 3})
	if err != nil {
		s.T().Fatal(err)
	}

	message, _ := s.store.FinishPlanting(&PlantedModule{Module: 3})

	expected := &Status{Message: "Planting Done"}
	assert.Equal(s.T(), message, expected)
}
