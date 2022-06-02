from flask_mongoengine import MongoEngine
from db import db_setup
from gpio import util
mongo = MongoEngine()

def create_app():
    

    from flask import Flask
    from .routes import routes

    app = Flask(__name__)
    app.config['MONGODB_SETTINGS'] = {
    'db': 'test',
    'host': 'mongodb',
    'port': 27017
    }
    
    #if util.isPi():
        #from gpio import cron
        #cron.scheduleAll
    try:
        mongo.init_app(app)
        db_setup.setup_db()
    except Exception as e: print(e)

    app.register_blueprint(routes)
    
    return app
    


import gpio
import db

