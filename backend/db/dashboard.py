from .schema import Module
from . import util


def get_harvestable_plants():
    harvestable_plants = []
    #first getting all plants that are older than age
    for module in Module.objects:
        if util.plants_in_module == 0:
            continue
        for plant in module.plants:
            print("smth")
            print(plant.variety.harvest_time)

            harvestable_plant = {
                "plant_type": planttype.name,
                "modulenum": module.modulenum
                #position

            }
            harvestable_plants.append(harvestable_plant)
    print(harvestable_plants)
    return harvestable_plants

def get_harvestable_leaves():
    return


def get_plantable_spots():
    #TODO add 7 day rule
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

        

