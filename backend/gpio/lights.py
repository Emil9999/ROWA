from gpiozero import LED
from . import pins

mainLightPin = LED(pins.pin.mainLight)

def mainLightToggle():
    mainLightPin.toggle()

def mainLightOn():
    mainLightPin.off()

def mainLightOff():
    mainLightPin.on()

def mainLightState():
    return not (mainLightPin.pin.state)
