from gpiozero import LED


mainLightPin = LED(6)
ledNumbersPin = 18

def mainLightToggle():
    mainLightPin.toggle()