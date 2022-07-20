from flask import Blueprint, request
from db import dashboard, harvesting, planting, admin_settings
import json
from gpio import lights, pumps

routes = Blueprint('routes', __name__)

@routes.route('/dashboard/farm/<module_number>')
def getPlantsInModule(module_number = 0):
    return json.dumps(dashboard.get_plants_in_module(module_number), default=str), 200, {'ContentType':'application/json'}


@routes.route("/plant", methods=['GET', 'POST'] )
def plant():
    if request.method == 'POST':
        if planting.insert_plant(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
       return json.dumps(dashboard.get_plantable_spots(), default=str), 200, {'ContentType':'application/json'} 




@routes.route("/harvest", methods=['GET', 'POST'] )
def harvest():
    if request.method == 'POST':
        if harvesting.harvest_plant(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return json.dumps(dashboard.get_harvestable_plants(), default=str), 200, {'ContentType':'application/json'} 
     

@routes.route("/admin/pump/times", methods=['GET', 'POST'] )
def pumpTimes():
    if request.method == 'POST':
        if admin_settings.insert_pump_times(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"

    else:
        if admin_settings.get_pump_times():
            return json.dumps(admin_settings.get_pump_times(), default=str), 200, {'ContentType':'application/json'}
        else:
            return "404"
   

@routes.route("/admin/pump")
@routes.route("/admin/pump/<state>")
def pump(state = None):
    if state == None:
        return pumps.pumpState()
    elif state == "on":
        pumps.pumpOn()
        return "True"
    elif state == "off":
        pumps.pumpOff()
        return "True"
    else:
        return "404"
    

@routes.route("/admin/airstone")
@routes.route("/admin/airstone/<state>")
def airstone(state = None):
    if state == None:
        return pumps.airstoneState()
    elif state == "on":
        pumps.airstoneOn()
        return "True"
    elif state == "off":
        pumps.airstoneOff()
        return "True"
    else:
        return "404"

@routes.route("/admin/light")
@routes.route("/admin/light/<state>")
def light(state = None):
    if state == None:
        return lights.lightState()
    elif state == "on":
        lights.mainLightOn()
        return "True"
    elif state == "off":
        lights.mainLightOff()
        return "True"
    else:
        return "404"

@routes.route("/admin/light/times", methods=['GET', 'POST'] )
def lightTimes():
    if request.method == 'GET':
        if admin_settings.get_light_times():
            return json.dumps(admin_settings.get_light_times(), default=str), 200, {'ContentType':'application/json'}
        else:
            return "404"
    elif request.method == 'POST':
        if admin_settings.insert_light_times(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"

@routes.route("/admin/planttypes")
def allVarieties():
    return json.dumps(admin_settings.get_all_varieties(), default=str), 200, {'ContentType':'application/json'}

@routes.route("/admin/change-planttype", methods=['GET', 'POST'] )
def changePlanttype():
    if request.method == 'POST':
        if admin_settings.change_planttype(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return "False"

@routes.route("/admin/reality-check", methods=['POST'] )
def realityCheck():
    if request.method == 'POST':
        if admin_settings.reality_check(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return "GET not supported"