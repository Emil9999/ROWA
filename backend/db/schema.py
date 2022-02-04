from mongoengine import *

#TODO add choices, default, required for the different fields
class Variety(Document):
    meta = {'collection': 'varieties'}
    name = StringField(primary_key=True)
    moving = BooleanField()
    harvest_time = IntField() #what do we set for herbs? init harvest?
    height = IntField() #definbes heights that make sense
    size = IntField() #what was it for?
    leaves_harvestable = IntField() #what was it for?
    harvests_per_week = IntField() # aka max harvests per week?

class Plant(EmbeddedDocument):
    variety = ReferenceField(Variety)
    plant_date = DateTimeField()
    user_name = StringField(default = "") #who planted it?
    position = IntField()
    leaves_harvested_total = IntField() #lifetime? 
    harvests_this_week = IntField(default = 0) #leaf harvests per week?
    last_harvest = DateTimeField(default = None)

class Module(Document):
    meta = {'collection': 'modules'}
    modulenum = IntField(min_value=1, max_value=50, primary_key=True)
    plant_spots = IntField(default = 6) #plant spots
    height = IntField(default = 1) #choice of 1 or 2 for now
    plants = ListField(EmbeddedDocumentField(Plant))

