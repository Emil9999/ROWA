//WaterTemp Libs
#include <OneWire.h>
#include <DallasTemperature.h>
//Adafruit Humidity and Temp
#include <Wire.h>
#include <Adafruit_AM2315.h>
//pH sensor
#include <SoftwareSerial.h>

//RC Emitter
#include <RCSwitch.h>

//In box Temp 
#include <math.h>
const int ntcWiderstand = 10000; 
const int MAX_ANALOG_VALUE = 1023;
#define PIN A2 

//Define for Emitter
RCSwitch sender = RCSwitch();

// Led Pins for ShiftReg
int datapin = 8; 
int clockpin = 10;
int latchpin = 9;

byte data = 0;


//sensor Pins
  //Legacy
const int tempPin = 0;
int photocellPin = 1;
  //UltraSonic
int trigPin = 5;
int echoPin = 6;
long duration;
  //WaterTemp
#define ONE_WIRE_BUS 4
OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);
 //HumidityAndTemp
Adafruit_AM2315 am2315;
struct temphum {
  float temp;
  float humi;
};
 //pH
 #define rx 2
 #define tx 3
SoftwareSerial myserial(rx, tx);
float pH;
String sensorstring = "";
boolean sensor_string_complete = false;

  
// Dynamic Variables for LED blink
int ledPin;
bool blinkLed;

bool ledOn = true;

// Variables for Sensors
int printPeriod = 1000 * 5; // every Minute
unsigned long time_now = 0;


void setup() {
  //LightSetup
  pinMode(datapin, OUTPUT);
  pinMode(clockpin, OUTPUT);  
  pinMode(latchpin, OUTPUT);

  

  //UltraSonic Setup
  pinMode(trigPin, OUTPUT);
  pinMode(echoPin, INPUT);

  //Emitter
  sender.enableTransmit(3);  // An Pin 3

  sender.setProtocol(1);
  sender.setPulseLength(300);



  //WaterTemp
  sensors.begin();

  //pH Node
  sensorstring.reserve(30);
  
  Serial.begin(9600);
  myserial.begin(9600);


   //Wait for HumTemp to start
  if (! am2315.begin()) {
     while (1);
  }
  
  while (!Serial) {
    ;
  }

  changeAll(false);
}

void loop() {
  // Wait on Serial Input to turn LED on
  if (Serial.available() > 0){
    int module_number = Serial.parseInt();
    switch (module_number){
      case 1:
        LedOn(0);
        break;
      case 2:
        LedOn(1);
        break;
      case 3:
       LedOn(2);
        break;
      case 4:
        LedOn(3);
        break;
      case 5:
       LedOn(4);
        break;
      case 6:
        LedOn(5);
        break;
      //Numbers for global off and on
      case 80:
       ledOn=true;
       changeAll(ledOn);
        break;
      case 81:
        ledOn=false;
        changeAll(ledOn);
        break;
      // Just a random number to signalise to turn off any led -> 99
      case 90:
        StartPump();
        break;
       case 91:
        StopPump();
        break;
       case 70:
        StartAir();
        break;
       case 71:
        StopAir();
        break;
      case 99:
        LedOff();
      default:
        break;
    }
    //Flushes remaining chars
    while (Serial.available() != 0 ) {
      Serial.read();
    }
  }


  if(millis() > time_now + printPeriod){
    time_now = millis();
    struct temphum temphum_instance = getHumTemp();
    int lightIntensity = getLightIntensity();
    long distanceToW = getDistance();
    float waterTemp = getWaterTemp();
    float waterpH = getpH();
    double boxTemp = getBoxTemp();
    Serial.println((String)temphum_instance.temp+","+(int)boxTemp+","+(String)temphum_instance.humi+","+(String)distanceToW+","+(String)waterTemp+","+(String)waterpH);

  }

  if (blinkLed){
    shiftWrite(ledPin, LOW);
    delay(500);
    shiftWrite(ledPin, HIGH);
    delay(500);
  }

}

//TestingFor phNode sending Stuff

void StartPump(){
  shiftWrite(6, HIGH);
}

void StopPump(){
  shiftWrite(6, LOW);
}

void StartAir(){
  shiftWrite(7, HIGH);
}

void StopAir(){
  shiftWrite(7, LOW);
}

void LedOn(int pin_number){
  ledPin = pin_number;
  blinkLed = true;
}

void LedOff(){
  blinkLed = false;
  if(ledOn){
    shiftWrite(ledPin, HIGH); }
    else{
      shiftWrite(ledPin, LOW);
    }
    }



int getLightIntensity(){
  int photocellReading = analogRead(photocellPin);
  return photocellReading;
}

long getDistance() {
  digitalWrite(trigPin, LOW);
  delayMicroseconds(5);
  digitalWrite(trigPin, HIGH);
  delayMicroseconds(10);
  digitalWrite(trigPin, LOW);

  pinMode(echoPin, INPUT);
  duration = pulseIn(echoPin, HIGH);

  //Returns distance in CM
  return (duration/2) / 29.1;
 }



 void changeAll(bool State)
{
// This function will turn on all the LEDs, one-by-one,
// and then turn them off all off, one-by-one. 

  int index;
  int delayTime = 100; // Time (milliseconds) to pause between LEDs
                       // Make this smaller for faster switching

 if(State){
  for(index = 0; index <= 5; index++)
  {
    bitWrite(data,index,HIGH); //Change desired bit to 0 or 1 in "data"

  // Now we'll actually send that data to the shift register.
  // The shiftOut() function does all the hard work of
  // manipulating the data and clock pins to move the data
  // into the shift register:

  
    delay(delayTime);                
  }
  shiftOut(datapin, clockpin, MSBFIRST, data);
  digitalWrite(latchpin, HIGH); 
  digitalWrite(latchpin, LOW); 
 }
 else{
  for(index = 0; index <= 5; index++)
  {
   bitWrite(data,index,LOW); //Change desired bit to 0 or 1 in "data"

  // Now we'll actually send that data to the shift register.
  // The shiftOut() function does all the hard work of
  // manipulating the data and clock pins to move the data
  // into the shift register:

  
    delay(delayTime);                
  }
  shiftOut(datapin, clockpin, MSBFIRST, data);
  digitalWrite(latchpin, HIGH); 
  digitalWrite(latchpin, LOW);
 }
 }



 
//Code for shift Reg
void shiftWrite(int desiredPin, boolean desiredState){

// This function lets you make the shift register outputs
// HIGH or LOW in exactly the same way that you use digitalWrite().

  bitWrite(data,desiredPin,desiredState); //Change desired bit to 0 or 1 in "data"

  // Now we'll actually send that data to the shift register.
  // The shiftOut() function does all the hard work of
  // manipulating the data and clock pins to move the data
  // into the shift register:

  shiftOut(datapin, clockpin, MSBFIRST, data); //Send "data" to the shift register

  //Toggle the latchPin to make "data" appear at the outputs
  digitalWrite(latchpin, HIGH); 
  digitalWrite(latchpin, LOW); 
}

double getBoxTemp(){
    float analogValue = analogRead(PIN);
  
  // Konvertieren des analogen Wertes in ein Widerstandswert
  float resistorValue = (MAX_ANALOG_VALUE / analogValue)- 1; 
  resistorValue = ntcWiderstand / resistorValue;
  double kelvin = convert2TempKelvin(analogValue);
  double celsius = convertKelvin2TempCelsius(kelvin);

  return celsius;
}

double convert2TempKelvin(float value){
  double temp = log(((10240000/value) - ntcWiderstand));
  temp = 1 / (0.001129148 + (0.000234125 * temp) + (0.0000000876741 * temp * temp * temp));
  return temp;
}
double convertKelvin2TempCelsius(double kelvin){
 return kelvin - 273.15;
}

float getWaterTemp(){
  sensors.requestTemperatures();
  return sensors.getTempCByIndex(0);
}

 struct temphum getHumTemp(){
  struct temphum temphum_instance;
  am2315.readTemperatureAndHumidity(&temphum_instance.temp, &temphum_instance.humi);
  return temphum_instance;
}


float getpH(){
 while(!sensor_string_complete){
  if (myserial.available() > 0) {                     //if we see that the Atlas Scientific product has sent a character
    char inchar = (char)myserial.read();              //get the char we just received
    sensorstring += inchar;                           //add the char to the var called sensorstring
    if (inchar == '\r') {                             //if the incoming character is a <CR>
      sensor_string_complete = true;                  //set the flag
    }
  }
 }
 
  if(sensor_string_complete == true){
    pH = sensorstring.toFloat();
  }
  sensorstring="";
  sensor_string_complete = false;
  return pH;
}
