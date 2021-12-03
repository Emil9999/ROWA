from flask_pymongo import PyMongo
from flask import g, current_app
from main import app

def get_db():
    if 'mongo' not in g:
        g.mongo = PyMongo(app)

    return g.mongo
