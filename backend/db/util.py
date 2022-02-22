from .schema import Plant, Module


def module_full(modulenum):
    module = Module.objects(modulenum = modulenum).get()
    module_full = Module.objects(modulenum = modulenum, plants__size = module.plant_spots).count()
    if module_full == 0:
        return False
    else:
        return True

def plants_in_module(modulenum):
    plantsinmodule = list(Module.objects(modulenum=modulenum).aggregate([{"$project": { "_id":0,"plant_count": { "$size":"$plants" }}}]))
    print(plantsinmodule[0]['plant_count'])
    return plantsinmodule[0]['plant_count']