from .schema import Module
from . import util


def get_harvestable_plants():
    harvestable_plants = []
    for module in Module.objects:
        if util.plants_in_module == 0:
            continue
        for planttype in module.plantable_varieties:

            plantable_plant = {
                "plant_type": planttype.name,
                "modulenum": module.modulenum
                #position

            }
            plantable_plants.append(plantable_plant)
    print(plantable_plants)   
    return harvestable_plants

def get_harvestable_leaves():
    return


def get_plantable_spots():
    plantable_plants = []
    for module in Module.objects:
        if util.module_full(module.modulenum):
            continue
        for planttype in module.plantable_varieties:

            plantable_plant = {
                "plant_type": planttype.name,
                "modulenum": module.modulenum
                #position

            }
            plantable_plants.append(plantable_plant)
    print(plantable_plants)   

        

