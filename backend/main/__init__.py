
from flask_pymongo import PyMongo


mongo = PyMongo()

def create_app():
    from flask import Flask
    from .routes import routes
    app = Flask(__name__)
    app.config["MONGO_URI"] = "mongodb://localhost:27017/test"
    app.register_blueprint(routes)
    mongo.init_app(app)
    return app


import db

