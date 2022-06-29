from tokenize import group
import datetime
from .schema import Plant, Module

def insert_plant(content):
    #plantsinmodule = list(Module.objects(modulenum=content['modulenum']).aggregate([{"$project": { "_id":0,"plant_count": { "$size":"$plants" }}}]))
    #print(plantsinmodule[0]['plant_count'])
    #implement checks if planttype matches, height, size, module not full etc
    #implement moving
    module = Module.objects(modulenum = content['module']).get()
    module_full = Module.objects(modulenum = content['module'], plants__size = module.plant_spots).count()
    if module_full == 0:
        #insert plant
        print('success')
        plant = Plant(
            variety = content['variety'],
            position = content['position'],
            planter = content['planter'],
                     )
        module.plants.append(plant)
        module.save()
        return True
    else:
        return False
   
