from .schema import Module, Variety, Settings

def change_planttype(content):
    module = Module(plantable_varieties= content['varieties'])
    updated_module = Module.objects(modulenum= content['modulenum']).modify(plantable_varieties = module.plantable_varieties)
    updated_module.save()                                                                      
    return


def insert_date(target, date):
    settings = Settings.objects.first()
    match target:
        case "pump":
            settings.update(set__pumpDate=date)
        case "light_on":
            settings.update(set__lightDateOn=date)
        case "light_off":
            settings.update(set__lightDateOff=date)
    