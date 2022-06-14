from flask import Blueprint, request
from db import dashboard, planting, admin_settings
import json

routes = Blueprint('routes', __name__)

@routes.route('/dashboard/farm/<module_number>')
def getPlantsInModule(module_number = 0):
    return json.dumps(dashboard.get_plants_in_module(module_number), default=str), 200, {'ContentType':'application/json'}


@routes.route("/plant/finish", methods=['GET', 'POST'] )
def plantFinish():
    if request.method == 'POST':
        if planting.insert_plant(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return "False"

@routes.route("/admin/change-planttype", methods=['GET', 'POST'] )
def changePlanttype():
    if request.method == 'POST':
        if admin_settings.change_planttype(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return "False"

@routes.route("/dashboard/harvestable-plants", defaults={'module_number': None})
@routes.route("/dashboard/harvestable-plants/<module_number>")
def getHarvestablePlants(module_number):
    #TODO Add module number as optional input to db func
    dashboard.get_harvestable_plants()
    return "True"

@routes.route("/dashboard/plantable-plants", defaults={'module_number': None})
@routes.route("/dashboard/plantable-plants/<module_number>")
def getPlantablePlants(module_number):
    #TODO Add module number as optional input to db func
    dashboard.get_plantable_spots()
    return "True"


#@routes.route("/dashboard/plantable-modules")

