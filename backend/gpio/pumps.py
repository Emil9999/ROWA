
from gpiozero import LED
from . import pins

pumpPin = LED(pins.pin.waterPump)
airPin =LED(pins.pin.airPump)
   
def pumpToggle():
    pumpPin.toggle()

def pumpOn():
    print("pumpon")
    pumpPin.on()

def pumpOff():
    print("pumpoff")
    pumpPin.off()

def airstoneToggle():
    airPin.toggle()

def airstoneOn():
    print("airstoneon")
    airPin.on()

def airstoneOff():
    print("airstoneoff")
    airPin.off()

def pumpState():
    return pumpPin.pin.state
 
def airstoneState():
    return airPin.pin.state