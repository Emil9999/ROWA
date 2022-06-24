from flask import Blueprint, request
from db import dashboard, planting, admin_settings
from gpio import util
import json

if util.isPi():
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
        return dashboard.get_plantable_spots()
        
@routes.route("/harvest", methods=['GET', 'POST'] )
def harvest():
    if request.method == 'POST':
        if planting.insert_plant(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return dashboard.get_harvestable_plants()
     


@routes.route("/admin/toggle-pump")
def togglePump():
    pumps.pumpToggle()
    return "True"

@routes.route("/admin/toggle-airstone")
def airstoneToggle():
    pumps.airstoneToggle()
    return "True"

@routes.route("/admin/toggle-light")
def toggleLight():
    lights.mainLightToggle()
    return "True"

@routes.route("/admin/change-planttype", methods=['GET', 'POST'] )
def changePlanttype():
    if request.method == 'POST':
        if admin_settings.change_planttype(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return "False"
