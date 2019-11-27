## Inroduction

###### ROWA project is consisted of two parts: Physical development of an indoor farm for offices and an application which is developed to interact with the farm. 

## Workflow Strategy

###### we decided to have a Feature Branch Workflow. This encapsulation makes it easy for multiple developers to work on a particular feature without disturbing the main codebase. It also means the master branch will never contain broken code, which is a huge advantage for continuous integration environments. Also, the Git Feature Branch Workflow is a composable workflow that can be leveraged by other high-level Git workflows.

## Software Architecture 

###### The application backend has been written in GO while the frontend is powered by using Vue.js framework and they run locally. Also, SQLite relational database management has been chosen for the database management system.

## Goal of the technical Documentation

###### The application have different features and I have decided to write a technical documentation for the plant feature and sensors part.

## Planting Feature
###### There are 2 main functions in the planting feature: 
    1. Plant 
###### When the user clicks on the planting button, he/she can choose a plant type. Afterwards, a post request will be sent to the server and the received data will be bounded to the plantType variable which stores the data in the json format. The database will be opened and the modules which have available spots and the same plant category will be selected and only the last position will be returned in json format since the user can plant in only one module. Accordingly, the function “ActivateModuleLight” in sensors.go file will run and the LED light of the shown position on the module will be turned on.

    2. FinishPlanting
###### The position of the plants will be changed one week up by the user when he wants to plant a new vegetable in the module. Therefore, the PlantPosition table will be updated by one while the data related to the new planted vegetable will be inserted into the Plant database table. At last, the function “DeactivateModuleLight” in sensors.go file will run and the LED light of the shown position on the module will be turned off.

###### Since I mentioned different database tables, I believe the database scheme would be helpful for the reader.

![alt text](https://github.com/MarcelCode/ROWA/blob/dev/documentation/Database%20Rowa.png)


