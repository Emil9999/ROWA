from flask_mongoengine import MongoEngine
from db import db_setup
from gpio import pins

mongo = MongoEngine()


def create_app():
    

    from flask import Flask
    from .routes import routes

    app = Flask(__name__)
    app.config['MONGODB_SETTINGS'] = {
    'db': 'test',
    'host': '127.0.0.1',
    'port': 27017
    }

    from gpio import cron
    cron.scheduleAll

    try:
        mongo.init_app(app)
        db_setup.setup_db()
    except Exception as e: print(e)

    app.register_blueprint(routes)
    print("Server ready..")
    return app
    


import gpio
import db

