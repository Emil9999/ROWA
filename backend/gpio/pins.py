import json
from types import SimpleNamespace

pin = None
def readPinsFromJson():
    try:
        with open('data/pins.json') as json_file:
            print("Reading pins from json")
            global pin
            pin = json.load(json_file, object_hook=lambda d: SimpleNamespace(**d))

    except FileNotFoundError:
        print("Json File not found")

