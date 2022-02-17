from .schema import Module


def get_harvestable_plants():
    print(harvestable_plants)
    return harvestable_plants

def get_plantable_spots():

    for module in Module:
        module = Module.objects(modulenum = module.modulenum).get()
        module_full = Module.objects(modulenum = module.modulenum, plants__size = module.plant_spots).count()

