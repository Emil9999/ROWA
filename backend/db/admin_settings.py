from .schema import Module, Variety

def change_planttype(content):
    module = Module(plantable_varieties= content['varieties'])
    updated_module = Module.objects(modulenum= content['modulenum']).modify(plantable_varieties = module.plantable_varieties)
    updated_module.save()                                                                      
    return