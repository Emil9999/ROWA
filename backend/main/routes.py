from flask import Blueprint, request
from db import dashboard, planting
import json

routes = Blueprint('routes', __name__)

@routes.route("/plant/finish", methods=['GET', 'POST'] )
def plantFinish():
    if request.method == 'POST':
        if planting.insert_plant(request.get_json()):
            return json.dumps({'success':True}), 200, {'ContentType':'application/json'} 
        else:
            return "404"
    else:
        return "False"
    
@routes.route("/dashboard/harvestable-plants")
def getHarvestablePlants():
    dashboard.get_harvestable_plants()
    return "True"

#@routes.route("/dashboard/plantable-plants")

#@routes.route("/dashboard/plantable-modules")

