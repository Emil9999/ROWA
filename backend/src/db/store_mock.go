package db

import (
	"github.com/stretchr/testify/mock"
)

// The mock store contains additonal methods for inspection
type MockStore struct {
	mock.Mock
}

func (m *MockStore) GetPlantsPerType(p string) ([]*PlantsPerPlantType, error) {
	rets := m.Called(p)
	return rets.Get(0).([]*PlantsPerPlantType), rets.Error(1)
}

func (m *MockStore) GetHarvestablePlant(plantType *PlantType) (*PositionOnFarm, error) {
	rets := m.Called(plantType)
	return rets.Get(0).(*PositionOnFarm), rets.Error(1)
}

func (m *MockStore) HarvestDone(plantPosition *PositionOnFarm) (*Status, error) {
	rets := m.Called(plantPosition)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) Plant(plantType *PlantType) (int, error) {
	rets := m.Called(plantType)
	return rets.Get(0).(int), rets.Error(1)
}

func (m *MockStore) FinishPlanting(plantedModule *PlantedModule) (*Status, error) {
	rets := m.Called(plantedModule)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) DbSetup() error {
	rets := m.Called()
	return rets.Error(0)
}

func (m *MockStore) GetLastSensorEntry() (sensorData *SensorData, err error) {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler tests)
	*/
	rets := m.Called()
	return rets.Get(0).(*SensorData), rets.Error(1)
}

func InitMockStore() *MockStore {
	/*
		Like the InitStore function we defined earlier, this function
		also initializes the store variable, but this time, it assigns
		a new MockStore instance to it, instead of an actual store
	*/
	s := new(MockStore)
	FunctionStore = s
	return s
}
