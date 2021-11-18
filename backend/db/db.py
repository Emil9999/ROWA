from flask_pymongo import PyMongo


_mongo = None

def get_connection(app=None):
    global _mongo
    if not _mongo:
        _mongo = PyMongo(app)
    return _mongo
