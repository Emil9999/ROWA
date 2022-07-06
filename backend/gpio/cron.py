
import time

from backend.main.routes import light
from . import pumps
from apscheduler.schedulers.background import BackgroundScheduler
from apscheduler.triggers.interval import IntervalTrigger
from db import admin_settings
from datetime import datetime, timedelta




def scheduleBoot():
    lightTimes = admin_settings.get_light_times()
    print(lightTimes['lightOn'].hour)
    pumpTimes = admin_settings.get_pump_times()
    pumpOffTime = pumpTimes['pumpDate'] + timedelta(minutes= pumpTimes['pumpDuration'])
    print("Scheduling jobs..")

    scheduler = BackgroundScheduler()
    scheduler.add_job(pumps.airstoneOn, 'cron', id='airOn',replace_existing=True,hour= pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute)
    scheduler.add_job(pumps.airstoneOff, 'cron',  id='airOff', replace_existing=True,hour = pumpOffTime.hour, minute= pumpOffTime.minute)
    scheduler.add_job(pumps.pumpOn, 'cron',  id='pumpOn', replace_existing=True,hour = pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute)
    scheduler.add_job(pumps.pumpOff, 'cron', id='pumpOff', replace_existing=True, hour = pumpOffTime.hour, minute= pumpOffTime.minute)

    scheduler.add_job(light.mainLightOn, 'cron',  id='lightOn',replace_existing=True,hour = lightTimes['lightOn'].hour, minute= lightTimes['lightOn'].minute)
    scheduler.add_job(light.mainLightOff, 'cron',  id='lightOff',replace_existing=True,hour = lightTimes['lightOff'].hour, minute= lightTimes['lightOff'].minute)

    scheduler.start()

def updateSchedule():
    lightTimes = admin_settings.get_light_times()
    print(lightTimes['lightOn'].hour)
    pumpTimes = admin_settings.get_pump_times()
    pumpOffTime = pumpTimes['pumpDate'] + timedelta(minutes= pumpTimes['pumpDuration'])

    print("Rescheduling jobs..")
    scheduler = BackgroundScheduler()
    scheduler.reschedule_job('airOn', hour= pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute)
    scheduler.reschedule_job('airOff', hour = pumpOffTime.hour, minute= pumpOffTime.minute)
    scheduler.reschedule_job('pumpOn', hour = pumpTimes['pumpDate'].hour, minute= pumpTimes['pumpDate'].minute)
    scheduler.reschedule_job('pumpOff', hour = pumpOffTime.hour, minute= pumpOffTime.minute)
    scheduler.reschedule_job('lightOn', hour = lightTimes['lightOn'].hour, minute= lightTimes['lightOn'].minute)
    scheduler.reschedule_job('lightOff', hour = lightTimes['lightOff'].hour, minute= lightTimes['lightOff'].minute)
