from mongoengine import *
import datetime

#TODO add choices, default, required for the different fields
class Variety(Document):
    meta = {'collection': 'varieties'}
    name = StringField(primary_key=True)
    moving = BooleanField()
    harvest_time = IntField() #what do we set for herbs? init harvest?
    height = IntField() #definbes heights that make sense
    size = IntField() #
    leaves_harvestable = IntField() #
    harvests_per_week = IntField() # aka max harvests per week?
    group = ListField(StringField()) #why listfield?

class Plant(EmbeddedDocument):
    variety = ReferenceField(Variety)
    plant_date = DateTimeField(default=datetime.datetime.utcnow)
    user_name = StringField(default = "") #who planted it?
    position = IntField(default = 1)
    leaves_harvested_total = IntField(default= 0) #lifetime
    harvests_this_week = IntField(default = 0) #leaf harvests per week?
    last_harvest = DateTimeField(default = None) #
    #array for harvest dates, then check array length, remove everything thats older than 7 weeks

class Module(Document):
    meta = {'collection': 'modules'}
    modulenum = IntField(min_value=1, max_value=50, primary_key=True)
    plant_spots = IntField()
    height = IntField(default = 1) #choice of 1 or 2 for now
    plants = ListField(EmbeddedDocumentField(Plant))
    plantable_varieties = ListField(ReferenceField(Variety))
