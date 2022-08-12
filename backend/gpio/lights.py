from gpiozero import LED
from . import pins

mainLightPin = LED(pins.pin.mainLight)
ledNumbersPin = 18

def mainLightToggle():
    mainLightPin.toggle()

def mainLightOn():
    mainLightPin.on()

def mainLightOff():
    mainLightPin.off()

def mainLightState():
    return mainLightPin.pin.state

