
import time
from . import pumps
from apscheduler.schedulers.background import BackgroundScheduler
from apscheduler.triggers.interval import IntervalTrigger


def scheduleAll():
    print("Scheduling jobs..")
    scheduler = BackgroundScheduler()
    
    scheduler.add_job(pumps.airstoneOn, 'cron', hour= 15, minute= 0)
    scheduler.add_job(pumps.airstoneOn, 'cron', second=20)
    scheduler.add_job(pumps.pumpOn, 'cron', second=10)
    scheduler.add_job(pumps.pumpOff, 'cron', second=20)

    scheduler.start()