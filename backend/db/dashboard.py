from .schema import Module
from . import util


def get_harvestable_plants():
    print(harvestable_plants)
    return harvestable_plants

def get_plantable_spots():
    plantable_plants = []
    for module in Module.objects:
        if util.module_full(module.modulenum):
            print(util.module_full(module.modulenum))
            continue
        for planttype in module.plantable_varieties:

            plantable_plant = {
                "plant_type": planttype,
                "modulenum": module.modulenum
                #position

            }
            plantable_plants.append(plantable_plant)
    print(plantable_plants)   

        

