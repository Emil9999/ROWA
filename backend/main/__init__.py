from flask import Flask, g
from flask.globals import current_app
from .routes import routes
from flask_pymongo import PyMongo


app = Flask(__name__)
app.config["MONGO_URI"] = "mongodb://localhost:27017/test"

mongo = PyMongo(app)

app.register_blueprint(routes)


import db

