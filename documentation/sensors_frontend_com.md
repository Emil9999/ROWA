# Technical documentation of data flow from the sensors to the frontend
## Introduction
In the ROWA project we are building an automated hydroponic indoor farm for offices. This documentation describes the way the data takes from being read from the sensor to being displayed on the screen of the farm. 
## Getting the sensor data to the screen
#### Reading the sensors:
We start by setting up the pins for our two sensors, then we  define a print period (in our case 1 minute). The reason for this is that the measurements don't change rapidly and thus making it unnecessary to read repetitive data. We also set up and open the serial port with a baud rate (or bit rate) of 9600 bit/s, which is used to send the data to our raspberry:
```c
Serial.begin(9600);
```
#### Sending the data to the raspberry:
For this we build a string that consists of our pH value and our conductivity value, seperated by a comma and followed by a new line.
```c
Serial.println((String)temp+","+lightIntensity);
```
#### Processing the data:
Moving on to the raspberry, we again set up and open a serial connection, of course with the same baud rate as the arduino. 
```go
c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
s, err = serial.OpenPort(c)
```
An imporant thing to know is that the data from the raspberry comes in a stream of bytes and not in one piece.
Thats why we have to create a buffer that fills up till it receives a new line, which marks the end of our string, as described above.
```go
buf := make([]byte, 128)
n, err := s.Read(buf)
```
Once we have the full string we split it into pH and conducvity respectively and write everything to our database.
```go
if strings.HasSuffix(serialString, "\n") {
  ph, err1 := strconv.ParseFloat(data_array[0], 32)
  ...
 }
```
That is it! We now simply have to display the data on the frontend using a simple GET request.
