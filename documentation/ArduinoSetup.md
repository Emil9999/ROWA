

Arduino Setup Representing the Physical Interface of the ROWA Farm
========

This documentation covers the circuit setup and the installation of the Arduino to represent the actual ROWA farm with 6
 LED stripes, one for each module and sensor measurements for the watertank (PH and Conductivity). Additionally it will
 be explained, how to control the Arduino over a serial connection via *Golang* functions.   

Background
-------

The ROWA farm is a small indoor farm where you can grow different kinds of plants like lettuce or herbs. With the attached
Screen you can choose, which plant you would like to plant or harvest. Additionally you can visualize sensor data to 
check the health of the plants.

The system consists of a Raspberry Pi, which handles a *Golang* webserver with a Sqlite Database and an
 Arduino which is responsible for the hardware interactions (Control lights and read sensor data). These two devices are
  communicating via serial communication with a baud rate of 9600 bit/s.

Circuit Setup
----------
![alt text](https://github.com/MarcelCode/ROWA/blob/master/documentation/images/circuit.png "Arduino Circuit")

Features
--------

- Control lights of farm modules via serial monitor or *Golang* functions
- Read sensor data via serial monitor or set up a goroutine which automatically adds sensor data to the database

Installation
------------

1. Download the Arduino Code [here](https://github.com/MarcelCode/ROWA/blob/dev/backend/src/sensor/sensors.go)
2. Connect Arduino to your computer via USB
2. Upload Code to Arduino Uno

Communication with the Arduino [optional]:
- Option 1:
    - Open serial monitor in Arduino Ide.
    - Sensor data can be seen once a minute.
    - Input Numbers **1-6** to turn on specific light. It is intended, that only one light at a time will shine.
    - Input Number **99** to turn off all lights.

- Option 2:
    - Control Arduino via *Golang* function.
    - It is recommended to start *ReadSensorData()* via goroutine, as it will automatically fill the database regularly. 
    
    ```go
    package sensors
  
    func ActivateModuleLight(moduleNumber int){} // Int 1-6
    
    func DeactivateModuleLight(){}
  
    func ReadSensorData(){}
    ```

Contribute
----------

- Issue Tracker: https://github.com/MarcelCode/ROWA/issues
- Source Code Arduino: https://github.com/MarcelCode/ROWA/tree/dev/arduino/arduino_uno
- Source Code *Golang* functions: https://github.com/MarcelCode/ROWA/blob/dev/backend/src/sensor/sensors.go

Support
-------

If you are having issues, please let us know.

