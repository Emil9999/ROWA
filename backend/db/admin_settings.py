from .schema import Module, Variety, Settings, Plant
from .util import string_to_datetime
import gpio, datetime

def change_planttype(content):
    module = Module(plantable_varieties= content['varieties'])
    module.plants.delete()
    updated_module = Module.objects(modulenum= content['modulenum']).modify(plantable_varieties = module.plantable_varieties)
    if updated_module.plantable_varieties[0].group[0] == 'herb':
        updated_module = Module.objects(modulenum= content['modulenum']).modify(plant_spots = 4)
    else:
        updated_module = Module.objects(modulenum= content['modulenum']).modify(plant_spots = 6)
    updated_module.save()                                                                      
    return

def get_pump_times():
    settings = Settings.objects.first()
    try:
        return {'onTime': settings.pumpDate,'duration': settings.pumpDuration}
    except AttributeError:
        return None

def get_light_times():
    settings = Settings.objects.first()
    try:
        return {'onTime': settings.lightDateOn, 'offTime': settings.lightDateOff}
    except AttributeError:
        return None

def insert_light_times(content):
    #TODO check that ontime is before offtime
    settings = Settings.objects.first()
    if settings == None:
        settings = Settings(lightDateOn=string_to_datetime(content['onTime']), lightDateOff=string_to_datetime(content['offTime']))
        settings.save()
    else:
        print(content)
        settings.update(set__lightDateOn=string_to_datetime(content['onTime']), set__lightDateOff=string_to_datetime(content['offTime']), upsert=True)
    gpio.cron.updateSchedule()
    return True

    
def insert_pump_times(content):
    print(content)
    settings = Settings.objects.first()
    if settings == None:
        settings = Settings(pumpDate=string_to_datetime(content['onTime']), pumpDuration=content['duration'])
        settings.save()
    else:
        settings.update(set__pumpDate=string_to_datetime(content['onTime']), set__pumpDuration=content['duration'], upsert=True)
    gpio.cron.updateSchedule()
    return True


def get_all_varieties():
    varieties = Variety.objects.all()
    vararray = []
    for variety in varieties:
        vararray.append({'plant_type': variety.name,'group': variety.group[0]})
    return vararray
    
        
def reality_check(content):
    module = Module.objects(modulenum= content[0]['module']).first()
    module.plants.delete()
    for x in range(len(content)):
        plant = Plant(variety=content[x]['plant_type'], position=content[x]['position'], plant_date= datetime.datetime.now()- datetime.timedelta(days=content[x]['age']))
        module.plants.append(plant)
    module.save()
    return True


def get_varieties_per_module(module):
    module = Module.objects(modulenum= module).first()
    varieties = module.plantable_varieties
    vararray = []
    for variety in varieties:
        vararray.append({'plant_type': variety.name,'group': variety.group[0], 'growth_time': variety.harvest_time})
    return vararray

def get_group_per_module(module):
    module = Module.objects(modulenum= module).first()
    varieties = module.plantable_varieties
    vararray = []
    for variety in varieties:
        print("vararray,", vararray)
        if not any(x for x in vararray if x == variety.group[0]):
            vararray.append(variety.group[0])
    
    return vararray