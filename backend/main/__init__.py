
from flask_mongoengine import MongoEngine
from db import db_setup

mongo = MongoEngine()

def create_app():
    from flask import Flask
    from .routes import routes

    app = Flask(__name__)
    app.config['MONGODB_SETTINGS'] = {
    'db': 'test',
    'host': 'localhost',
    'port': 27017
    }

    mongo.init_app(app)
    db_setup.setup_db()
    app.register_blueprint(routes)
    return app



import db

