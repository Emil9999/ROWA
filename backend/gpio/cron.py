
import time
from . import pumps, lights
from apscheduler.schedulers.background import BackgroundScheduler
from apscheduler.triggers.interval import IntervalTrigger
import db
from datetime import date, datetime, timedelta


scheduler = BackgroundScheduler()

def scheduleBoot():
    lightTimes = db.admin_settings.get_light_times()
    print(lightTimes['onTime'].hour)
    pumpTimes = db.admin_settings.get_pump_times()
    pumpOffTime = pumpTimes['onTime'] + timedelta(minutes= pumpTimes['duration'])
    print("Scheduling jobs..")
    #Interval jobs
    scheduler.add_job(pumps.pumpOn, 'cron', id='pumpOn',replace_existing=True, hour= '8-18', minute='*/05', max_instances=1 ,timezone='Europe/Berlin')
    scheduler.add_job(pumps.airstoneOn, 'cron', id='airOn',replace_existing=True, hour= '8-18', minute='*/05', max_instances=1 ,timezone='Europe/Berlin')
    scheduler.add_job(pumps.pumpOff, 'cron', id='pumpOff',replace_existing=True, hour= '8-18', minute='*/15', max_instances=1 ,timezone='Europe/Berlin')
    scheduler.add_job(pumps.airstoneOff, 'cron', id='airOff',replace_existing=True, hour= '8-18', minute='*/15', max_instances=1 ,timezone='Europe/Berlin')



    #Normal cron jobs
    #scheduler.add_job(pumps.airstoneOn, 'cron', id='airOn',replace_existing=True,hour= pumpTimes['onTime'].hour, minute= pumpTimes['onTime'].minute, max_instances=1)
    #scheduler.add_job(pumps.airstoneOff, 'cron',  id='airOff', replace_existing=True,hour = pumpOffTime.hour, minute= pumpOffTime.minute , max_instances=1)
    #scheduler.add_job(pumps.pumpOn, 'cron',  id='pumpOn', replace_existing=True,hour = pumpTimes['onTime'].hour, minute= pumpTimes['onTime'].minute, max_instances=1)
    #scheduler.add_job(pumps.pumpOff, 'cron', id='pumpOff', replace_existing=True, hour = pumpOffTime.hour, minute= pumpOffTime.minute, max_instances=1)

    scheduler.add_job(lights.mainLightOn, 'cron',  id='lightOn',replace_existing=True,hour = lightTimes['onTime'].hour, minute= lightTimes['onTime'].minute, max_instances=1)
    scheduler.add_job(lights.mainLightOff, 'cron',  id='lightOff',replace_existing=True,hour = lightTimes['offTime'].hour, minute= lightTimes['offTime'].minute, max_instances=1)

    scheduler.start()

def updateSchedule():
    lightTimes = db.admin_settings.get_light_times()
    print(lightTimes['onTime'].hour)
    pumpTimes = db.admin_settings.get_pump_times()
    pumpOffTime = pumpTimes['onTime'] + timedelta(minutes= pumpTimes['duration'])

    print("Rescheduling jobs..")
    #scheduler.reschedule_job('airOn',trigger='cron', hour = pumpTimes['onTime'].hour, minute= pumpTimes['onTime'].minute)
    #scheduler.reschedule_job('airOff',trigger='cron',hour = pumpOffTime.hour, minute= pumpOffTime.minute)
    #scheduler.reschedule_job('pumpOn',trigger='cron', hour = pumpTimes['onTime'].hour, minute= pumpTimes['onTime'].minute)
    #scheduler.reschedule_job('pumpOff', trigger='cron',hour = pumpOffTime.hour, minute= pumpOffTime.minute)
    scheduler.reschedule_job('lightOn', trigger='cron',hour = lightTimes['onTime'].hour, minute= lightTimes['onTime'].minute)
    scheduler.reschedule_job('lightOff', trigger='cron',hour = lightTimes['offTime'].hour, minute= lightTimes['offTime'].minute)
