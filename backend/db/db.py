from flask_pymongo import PyMongo

_mongo = None

def get_connection(app=None):
    global _mongo
    if app:
        _mongo = PyMongo(app)
    else:
        return _mongo
