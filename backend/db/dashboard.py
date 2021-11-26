from flask_pymongo import PyMongo
from db import db

mongo = db.get_connection()

def get_harvestable_plants():
    basil = mongo.db.plants.find_one({"plantType": "Basil"})
    print(basil)
    return basil