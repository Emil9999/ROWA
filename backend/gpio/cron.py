
import time
from . import pumps, lights
from apscheduler.schedulers.background import BackgroundScheduler
from apscheduler.triggers.interval import IntervalTrigger
import db
from datetime import datetime, timedelta


scheduler = BackgroundScheduler()


def scheduleBoot():
    lightTimes = db.admin_settings.get_light_times()
    print(lightTimes['lightOn'].hour)
    pumpTimes = db.admin_settings.get_pump_times()
    pumpOffTime = pumpTimes['pumpDate'] + timedelta(minutes= pumpTimes['pumpDuration'])
    print("Scheduling jobs..")

    scheduler.add_job(pumps.airstoneOn, 'cron', id='airOn',replace_existing=True,hour= pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute, max_instances=1)
    scheduler.add_job(pumps.airstoneOff, 'cron',  id='airOff', replace_existing=True,hour = pumpOffTime.hour, minute= pumpOffTime.minute , max_instances=1)
    scheduler.add_job(pumps.pumpOn, 'cron',  id='pumpOn', replace_existing=True,hour = pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute, max_instances=1)
    scheduler.add_job(pumps.pumpOff, 'cron', id='pumpOff', replace_existing=True, hour = pumpOffTime.hour, minute= pumpOffTime.minute, max_instances=1)

    scheduler.add_job(lights.mainLightOn, 'cron',  id='lightOn',replace_existing=True,hour = lightTimes['lightOn'].hour, minute= lightTimes['lightOn'].minute, max_instances=1)
    scheduler.add_job(lights.mainLightOff, 'cron',  id='lightOff',replace_existing=True,hour = lightTimes['lightOff'].hour, minute= lightTimes['lightOff'].minute, max_instances=1)

    scheduler.start()

def updateSchedule():
    lightTimes = db.admin_settings.get_light_times()
    print(lightTimes['lightOn'].hour)
    pumpTimes = db.admin_settings.get_pump_times()
    pumpOffTime = pumpTimes['pumpDate'] + timedelta(minutes= pumpTimes['pumpDuration'])

    print("Rescheduling jobs..")
    scheduler.reschedule_job('airOn',trigger='cron', hour = pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute)
    scheduler.reschedule_job('airOff',trigger='cron',hour = pumpOffTime.hour, minute= pumpOffTime.minute)
    scheduler.reschedule_job('pumpOn',trigger='cron', hour = pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute)
    scheduler.reschedule_job('pumpOff', trigger='cron',hour = pumpOffTime.hour, minute= pumpOffTime.minute)
    scheduler.reschedule_job('lightOn', trigger='cron',hour = lightTimes['lightOn'].hour, minute= lightTimes['lightOn'].minute)
    scheduler.reschedule_job('lightOff', trigger='cron',hour = lightTimes['lightOff'].hour, minute= lightTimes['lightOff'].minute)
