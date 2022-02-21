from .schema import Plant, Module


def module_full(modulenum):
    module = Module.objects(modulenum = modulenum).get()
    module_full = Module.objects(modulenum = modulenum, plants__size = module.plant_spots).count()
    if module_full == 0:
        return False
    else:
        return True