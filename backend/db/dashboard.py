from db.db import get_db
from flask import current_app

#mongo = get_db()

def get_harvestable_plants():
    basil = mongo.db.plants.find_one({"plantType": "Basil"})
    print(basil)
    return basil