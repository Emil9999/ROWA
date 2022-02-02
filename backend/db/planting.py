from main import mongo
from .schema import Plant

def insert_plant(modulenum):
    module = mongo.db.plants.find_one_and_update({"modulenum": modulenum})
    #implement checks if planttype matches, height, size, module not full etc
    plant = Plant()
    #print(harvestable_plants)
    return "yay"