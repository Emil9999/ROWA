# Susteyn
 This is the official repo for the software that runs on the susteyn office farms. All software works without an internet connection and is running locally, including the database. Automatic OTA updates do require a connection but are optional.
 
## Components
#### Base
Currently our system is deployed on a raspberry pi 4, running RPI OS. All software components run in docker containers and are orchestrated using docker compose, see docker-compose.yaml for details. Builds are triggered via github actions and then pushed to docker hub. Watchtower then automatically fetches the newest containers and gracefully replaces them.

#### Frontend
Makes interacting with the farm possible. Written in vue 3. 

#### Backend
Handles frontend requests, reads sensor data and commands actuators, specificallly turning on and off lights and pumps. Written in pythin using flask.
#### Database
Local mongodb instance also running in containers.

## API
See the current API documentation [here.](https://emil9999.github.io/ROWA/#/)

## Deprecated features 
The telegraf and influxdb integrations are no longer in use because they were simply not needed.

## Running
Simply clone the dev branch and run `docker-compose -f docker-compose.dev.yaml up` (requires docker installed)

## Older farms
Older versions of the farm (2.0) were running a similar stack, however used vue 2 instead of 3, golang instead of flask and sqlite .instead of mongo. The current software is not backwards compatible and should only be used on 3.0 versions of the farm.
