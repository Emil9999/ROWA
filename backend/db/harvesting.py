from operator import mod
from .schema import Plant, Module
import datetime

def harvest_plant(content):
    #TODO need position and module number and verify this, also implement leaf harvest
    module = Module.objects(modulenum = content['module']).get()
    if content['leaf_harvest'] == True:
        plant = module.plants.get(position = content['position'])
        plant.leaves_harvested_total += 1
        plant.harvests_this_week += 1
        plant.harvest_dates.append(datetime.datetime.now())
        # what else to add here
        module.save()
        return True
    else:
        module.update(pull__plants__position = content['position'])
        module.save()
        
        return True
        