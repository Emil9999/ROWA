Getting the sensor data to the screen

Reading the sensors:

The farm collects two different types of sensor data: pH and conductivity. For this an arduino is used, which functions as a middlemand between the sensors and the
raspberry PI on which the fronted runs.
We start by setting up the pins for our two sensors, then we  define a print period (in our case 1 minute). This means that the arduino only reads and sends the sensor data only every minute, and not whenver the sensor readings update.
The reason for this is that the measurements don't change rapidly and thus making it unecessary to read repetitive data.
We also set up and open the serial port with a baud rate (or bit rate) of 9600 bit/s, which is used to send the data to our raspberry.
For this we build a string that consists of our pH value and our conductivity value, seperated by a comma and followed by a new line. This is all the arduino has to do.
Moving on to the raspberry, we again set up and open a serial connection, of course with the same baud rate as the arduino. We also open our sqlite database where we will save
all the data for not only displaying it, but also to be able to use it in the future. 
An imporant thing to know is that the data from the raspberry comes in a stream of bytes and not in one piece.
Thats why we have to create a buffer that fills up till it receives a new line, which marks the end of our string, as described above.
Once we have the full string we split it into pH and conducvity respectively, add a time of measurement and write everything to our database.
We then simply have to query our database via the frontend and TADA, we can see the sensor data on our beautiful display.
