# Technical documentation of data flow from the sensors to the frontend
## Introduction
In the ROWA project we are building an automated hydroponic indoor farm for offices. From a developers perspective, the farm consists of two parts: The physical/hardware aspect, which includes the frame, the plant pots and especially sensors and lights. On the other hand we have the software part, which ensures the correct data collection and storage, as well as the correct digital farm representation on the farm. It is crucial for the product to have a clean and neatless integration of both these parts to ensure the correct functionning of the farm, hence this documentation describes the way the data takes from being read from the sensor to being displayed on the screen of the farm. 

## Getting the sensor data to the screen

#### Reading the sensors:

The farm collects two different types of sensor data: pH and conductivity. For this an arduino is used, which functions as a middlemand between the sensors and the raspberry PI on which the backend runs locally.
We start by setting up the pins for our two sensors, then we  define a print period (in our case 1 minute). This means that the arduino only reads and sends the sensor data only every minute, and not whenever the sensor readings update:
```c
int tempPin = 0;
int photocellPin = 1;
int printPeriod = 5000;
```
The reason for this is that the measurements don't change rapidly and thus making it unnecessary to read repetitive data.
We also set up and open the serial port with a baud rate (or bit rate) of 9600 bit/s, which is used to send the data to our raspberry:
```c
Serial.begin(9600);
```
#### Sending the data to the raspberry:
For this we build a string that consists of our pH value and our conductivity value, seperated by a comma and followed by a new line. This is all the arduino has to do.
```c
Serial.println((String)temp+","+lightIntensity);
```
#### Processing the data:
Moving on to the raspberry, we again set up and open a serial connection, of course with the same baud rate as the arduino. 
Keep in mind that this now is GO code, since our backend is written completely in GO:
```go
c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
s, err = serial.OpenPort(c)
```

We also open our sqlite database where we will save all the data for not only displaying it, but also to be able to use it in the future. 
An imporant thing to know is that the data from the raspberry comes in a stream of bytes and not in one piece.
Thats why we have to create a buffer that fills up till it receives a new line, which marks the end of our string, as described above.
```go
buf := make([]byte, 128)
n, err := s.Read(buf)
```
Once we have the full string we split it into pH and conducvity respectively, add a time of measurement and write everything to our database.
```go
if strings.HasSuffix(serialString, "\n") {
  ph, err1 := strconv.ParseFloat(data_array[0], 32)
  ...
 }
```
#### Displaying the data:
We now simply have to send a GET request to the backend requesting the most recent sensor data from the DB.
The backend queries the DB and sends back the data in a JSON format and TADA, we can the sensor data on the screen.
