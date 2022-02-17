from .schema import Module


def get_harvestable_plants():
    print(harvestable_plants)
    return harvestable_plants

def get_plantable_spots():

    for module in Module:
        plantsinmodule = list(Module.objects(modulenum=module.modulenum).aggregate([{"$project": { "_id":0,"plant_count": { "$size":"$plants" }}}]))
        print(plantsinmodule[0]['plant_count'])
        

