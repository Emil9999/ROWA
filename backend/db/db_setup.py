import json
from .schema import Variety, Module, Settings


def insert_varieties(varieties):
    #Deleting all entries in varieties and rewriting on every startup 
    Variety.objects().delete()
    for i in varieties:
        variety = Variety(
            name=i['name'],
            moving=i['moving'],
            harvest_time=i['harvest_time'],
            height=i['height'],
            size=i['size'],
            leaves_harvestable=i['leaves_harvestable'],
            harvests_per_week=i['harvests_per_week'],
            #TODO check if group is is in variety groups, otherwise json faulty
            group=i['group']
        )
        variety.save()


def insert_modules(modules):
    #Updating existing or creating new modules
    for i in modules:
        plantable_varieties = []
        for variety in i['varieties']:
            plantable_varieties.append(Variety.objects(name = variety).get())
        module = Module(
            modulenum=i['modulenum'],
            plant_spots=i['size'],
            height=i['height'],
            plantable_varieties= plantable_varieties
        )
        #print(module.plantable_varieties)
        updated_module = Module.objects(modulenum=module.modulenum).modify(upsert=True, new=True,
                                                                           modulenum=module.modulenum,
                                                                           plant_spots= module.plant_spots,
                                                                           height=module.height,
                                                                           plantable_varieties = module.plantable_varieties
                                                                           )
        updated_module.save()

    # Deleting dropped modules
    Module.objects(modulenum__gt=len(modules)).delete()



def setup_db():
    print("Setting up db")
    try:
        with open('data/farm_data.json') as json_file:
            farm_data = json.load(json_file)
            insert_varieties(farm_data['varieties'])
            insert_modules( farm_data['modules'])

    except FileNotFoundError:
        print("Json File not found")

    
