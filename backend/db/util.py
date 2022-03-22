from .schema import Plant, Module


def module_full(modulenum):
    module = Module.objects(modulenum = modulenum).get()
    if plants_in_module(modulenum) >= module.plant_spots:
        return True
    else:
        return False

def plants_in_module(modulenum):
    return Module.objects.get(modulenum = modulenum).plants.count()