from gpiozero import Button
from . import pins

floaterLow = Button(pins.pin.floaterLow)
floaterHigh = Button(pins.pin.floaterHigh)
floaterMiddle = Button(pins.pin.floaterMiddle)
