from .schema import Module, Variety, Settings
from .util import string_to_datetime
from gpio import cron

def change_planttype(content):
    module = Module(plantable_varieties= content['varieties'])
    updated_module = Module.objects(modulenum= content['modulenum']).modify(plantable_varieties = module.plantable_varieties)
    updated_module.save()                                                                      
    return

def get_pump_times():
    settings = Settings.objects.first()
    try:
        return {'pumpDate': settings.pumpDate,'pumpDuration': settings.pumpDuration}
    except AttributeError:
        return None

def get_light_times():
    settings = Settings.objects.first()
    try:
        return {'lightOn': settings.lightDateOn, 'lightOff': settings.lightDateOff}
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
    cron.updateSchedule()
    return True

    
def insert_pump_times(content):
    print(content)
    settings = Settings.objects.first()
    if settings == None:
        settings = Settings(pumpDate=string_to_datetime(content['onTime']), pumpDuration=content['duration'])
        settings.save()
    else:
        settings.update(set__pumpDate=string_to_datetime(content['onTime']), set__pumpDuration=content['duration'], upsert=True)
    cron.updateSchedule()
    return True