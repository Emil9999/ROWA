from flask import Blueprint, current_app
from db import dashboard

routes = Blueprint('routes', __name__)


@routes.route("/dashboard/harvestable-plants")
def getHarvestablePlants():
    dashboard.get_harvestable_plants()
    return 