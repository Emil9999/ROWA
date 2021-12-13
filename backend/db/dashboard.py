from flask import current_app
from main import mongo

#mongo = get_db()

def get_harvestable_plants():
    harvestable_plants = mongo.db.plants.find_one({"plantType": "Basil"})
    print(harvestable_plants)
    return harvestable_plants