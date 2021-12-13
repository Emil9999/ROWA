from flask import Flask, g
from flask.globals import current_app
from .routes import routes
from flask_pymongo import PyMongo


app = Flask(__name__)
app.config["MONGO_URI"] = "mongodb://localhost:27017/test"
#app.app_context().push()
#app.config["DB"] = PyMongo(app)
mongo = PyMongo(app)




#Removing db connection from globals
#app.teardown_appcontext
#def teardown_db(exception):
  #  mongo = g.pop('mongo', None)



app.register_blueprint(routes)

# def get_db():
#     if 'mongo' not in g:
#         g.mongo = PyMongo(current_app)

#     return g.mongo



import db

