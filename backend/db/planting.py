from main import mongo
from .schema import Plant, Module

def insert_plant(content):
    #print(content['modulenum'])
    #print(Module.objects(modulenum = content['modulenum'], plantable_varieties = content['variety']))
    #module = mongo.db.plants.find_one_and_update({"modulenum": modulenum})
    #implement checks if planttype matches, height, size, module not full etc
    module = Module.objects(modulenum = content['modulenum']).get()
    print(module.plant_spots)
    # plants_inmobule = Module.objects(modulenum = content['modulenum'], plants__size = )
    # if Module.objects(modulenum = content['modulenum'], plantable_varieties = content['variety']) and Module.objects(modulenum = content['modulenum'], plants__size = module.plant_spots):
    #     #insert plant
    #     print('success')
    #     return True
    # else:
    #     return False
    # plant = Plant()
