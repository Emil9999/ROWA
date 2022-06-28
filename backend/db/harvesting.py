from .schema import Plant, Module

def harvest_plant(content):
    #TODO need position and module number and verify this
    module = Module.objects(modulenum = content['modulenum']).get()
    plant = Module.plants.objects(position = content['position']).get()
    plant.delete()
    return True