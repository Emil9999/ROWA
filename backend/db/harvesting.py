from .schema import Plant, Module

def harvest_plant(content):
    #TODO need position and module number and verify this, also implement leaf harvest
    module = Module.objects(modulenum = content['modulenum']).get()
    if content['leaf_harvest'] == True:
        plant = Plant.objects.get(position = content['position'], modulenum = content['modulenum'])
        plant.leaves_harvested_total += 1
        # what else to add here
        plant.save()
    else:
        plant = module.plants.objects(position = content['position']).get()
        plant.delete()

    return True