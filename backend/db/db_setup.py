from main import mongo
import json
from .schema import Variety

def insert_varieties():
    with open('data/varieties.json') as json_file:
            varieties = json.load(json_file)
            for i in varieties['varieties']:
                variety = Variety(
                    name = i['name'],
                    moving = i['moving'],
                    harvest_time = i['harvest_time'],
                    height = i['height'],
                    size = i['size'],
                    leaves_harvestable = i['leaves_harvestable'],
                    harvests_per_week = i['harvests_per_week']
                )
                print(variety)
                variety.update(upsert=True) #insert if not exist or update




def setup_db():
    insert_varieties()
