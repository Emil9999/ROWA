from flask import Blueprint, current_app, request
from db import dashboard, planting

routes = Blueprint('routes', __name__)

@routes.route("/plant/finish", methods=['GET', 'POST'] )
def plantFinish():
     if request.method == 'POST':
         planting.plant(request.get_json()),
         return "True"


@routes.route("/dashboard/harvestable-plants")
def getHarvestablePlants():
    dashboard.get_harvestable_plants()
    return "True"

#@routes.route("/dashboard/plantable-plants")

#@routes.route("/dashboard/plantable-modules")

