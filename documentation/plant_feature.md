## Inroduction

##### ROWA project is consisted of two parts: Physical development of an indoor farm for offices and an application which is developed to interact with the farm. 

## Workflow Strategy

##### we decided to have a Feature Branch Workflow. This encapsulation makes it easy for multiple developers to work on a particular feature without disturbing the main codebase. Also, the Git Feature Branch Workflow is a composable workflow that can be leveraged by other high-level Git workflows.

## Software Architecture 

##### The application backend has been written in GO while the frontend is powered by using Vue.js framework and they run locally. Also, SQLite relational database management has been chosen for the database management system.

## Planting Feature
##### There are 2 main functions in the planting feature: 
1. Plant & PlantHandler
##### The Plant function helps the user to plant a vegetable by sending a POST request containing the PlantType variable.
``` 
func Plant{
    plantType := new(utils.PlantType)
    ...
} 
``` 
##### The handler for this post request is PlantHandler and it takes the data which is in json format to pass it to the Plant function and the available module will be shown by the light which is triggered by the ActivateModuleLight function in sensors.go file.
``` 
func PlantHandler{
    ...
    return c.JSON(http.StatusOK, position)
}
```
 2. FinishPlanting & FinishPlantingHandler
##### The FinishPlanting function ends the planting action by increasing the position of the plants one week up after the user plant a new vegetable in the module. 
``` 
func FinishPlanting{
    plantedModule := new(PlantedModule)
    ...
}
```
##### Also, the handler for this function, FinishPlantingHandler, takes the json format data(PlabtedModule) to pass it to the FinishPlanting function and a new plant will be added to the allocated position in the module. 
``` 
func FinishPlantingHandler{
    ...
    return c.JSON(http.StatusOK, true)
}
```
##### Further, the module light which is triggered by the DeactivateModuleLight function in sensors.go file will be turned off.






