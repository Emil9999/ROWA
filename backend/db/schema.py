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
    leaves_harvestable = IntField(default=0) #
    harvests_per_week = IntField() # aka max harvests per week?
    group = ListField(StringField()) #why listfield?

class Plant(EmbeddedDocument):
    variety = ReferenceField(Variety)
    plant_date = DateTimeField(default=datetime.datetime.utcnow)
    planter = StringField(default = "") #who planted it
    position = IntField(default = 1)
    leaves_harvested_total = IntField(default= 0) #lifetime
    harvests_this_week = IntField(default = 0) #leaf harvests per week?
    harvest_dates = ListField(DateTimeField(default = None)) #
    #array for harvest dates, then check array length, remove everything thats older than 7 days

class Module(Document):
    meta = {'collection': 'modules'}
    modulenum = IntField(min_value=1, max_value=50, primary_key=True)
    plant_spots = IntField()
    height = IntField(default = 1) #choice of 1 or 2 for now
    plants = EmbeddedDocumentListField(Plant)
    plantable_varieties = ListField(ReferenceField(Variety))

class Settings(Document):
    meta = {'collection': 'settings'}
    pumpDateOn = DateTimeField(default = datetime.datetime.strptime('08:00', '%H:%M'))
    pumpDateOff = DateTimeField(default = datetime.datetime.strptime('18:00', '%H:%M'))
    pumpDuration = IntField(default = 30)
    lightDateOn = DateTimeField(default = datetime.datetime.strptime('05:00', '%H:%M'))
    lightDateOff = DateTimeField(default = datetime.datetime.strptime('22:00', '%H:%M'))    
