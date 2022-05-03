from gpiozero import LED

pumpPin = LED(17)
airPin =LED(27)
   

def pumpOn():
    print("pumpon")
    pumpPin.on()

def pumpOff():
    print("pumpoff")
    pumpPin.off()

def airstoneOn():
    print("airstoneon")
    airPin.on()

def airstoneOff():
    print("airstoneoff")
    airPin.off()

