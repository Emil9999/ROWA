from .schema import Module
import datetime
from . import util
from mongoengine.queryset.visitor import Q


def get_plants_in_module(module):
    plants_in_module = []
    
    for plant in Module.objects.get(modulenum = module).plants:
        plant = {
            "plant_type": plant.variety.name,
            "age": (datetime.datetime.now() - plant.plant_date).days ,
            "position": plant.position,
            "growth_time": plant.variety.harvest_time
        }
        plants_in_module.append(plant)
    return plants_in_module

def get_harvestable_plants():
    harvestable_plants = []
    # first getting all plants that are older than age
    for module in Module.objects:
        if util.plants_in_module == 0:
            continue
        for plant in module.plants:
            print(plant.plant_date +
                  datetime.timedelta(days=plant.variety.harvest_time))
            if plant.plant_date + datetime.timedelta(days=plant.variety.harvest_time) <= datetime.datetime.today():
                print("harvestable")
                harvestable_plant = {
                    "plant_type": plant.variety.name,
                    "module": module.modulenum,
                    "position": plant.position,
                    "group": plant.variety.group[0],
                    "leaf_harvest":leaves_harvestable(plant),
                    "planter": plant.planter,

                }
                harvestable_plants.append(harvestable_plant)
    print(harvestable_plants)
    return harvestable_plants

def leaves_harvestable(plant):
    if plant.variety.leaves_harvestable == 0:
        return False
    else:
        return True
        
def get_harvestable_leaves():
    return


def get_plantable_spots():
    # TODO test 7 day rule
    plantable_plants = []
    for module in Module.objects:
        if not seven_day_rule(module.modulenum):
            continue
        if util.module_full(module.modulenum):
            continue
        for planttype in module.plantable_varieties:
            #filter out herbs, not plantable by user
            if planttype.group == ['herb']:
                continue
            plantable_plant = {
                "plant_type": planttype.name,
                "module": module.modulenum,
                "position": plantable_spot(module),
                "group": planttype.group[0]
            }
            plantable_plants.append(plantable_plant)
            
    print(plantable_plants)
    return plantable_plants

def plantable_spot(module):
    #find all plantable spots
    spots = []
    for plant in module.plants:
        spots.append(plant.position)
    pspots = []
    for i in range(1, module.plant_spots):
        if not (i in spots):
            pspots.append(i)

    pspots.sort()
    return pspots[0]

    
def seven_day_rule(modulenum):
    for plant in Module.objects.get(modulenum=modulenum).plants:
        if plant.plant_date > datetime.datetime.today() - datetime.timedelta(days=7):
            return False

    return True
