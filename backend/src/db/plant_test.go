package db

import (
	"fmt"
)

func (s *StoreSuite) TestPlant() {
	modulePosition, err := s.store.Plant(&PlantType{Name: "Basil"})
	if err != nil {
		s.T().Fatal(err)
	}

	fmt.Println(modulePosition)
}
