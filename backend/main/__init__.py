from flask_mongoengine import MongoEngine
from db import db_setup
import os

mongo = MongoEngine()


def create_app():
    

    from flask import Flask
    from flask_cors import CORS, cross_origin
    from .routes import routes

    app = Flask(__name__)
    cors = CORS(app)
    app.config['CORS_HEADERS'] = 'Content-Type'
    app.config['MONGODB_SETTINGS'] = {
    'db': 'test',
    'host': "127.0.0.1",
    'port': 27017
    }

  

    try:
        mongo.init_app(app)
        db_setup.setup_db()
    except Exception as e: print(e)

    app.register_blueprint(routes)
    from gpio import cron
    cron.scheduleBoot()
    print("Server ready..")
    return app
    


import gpio
import db

