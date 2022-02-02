from mongoengine import *


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
    user_name = StringField() #who planted it?
    position = IntField()
    leaves_harvested_total = IntField() #lifetime? 
    harvests_this_week = IntField() #leaf harvests per week?
    last_harvest = DateTimeField()

class Module(Document):
    meta = {'collection': 'modules'}
    modulenum = IntField(min_value=1, max_value=6)
    size = IntField()
    height = IntField()
    plants = ListField(EmbeddedDocumentField(Plant))

