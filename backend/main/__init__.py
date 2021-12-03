from flask import Flask
from main.routes import routes


app = Flask(__name__)
app.config["MONGO_URI"] = "mongodb://localhost:27017/test"
#app.app_context().push()
#app.config["DB"] = PyMongo(app)
#mongo = PyMongo(app)





#Removing db connection from globals
#app.teardown_appcontext
#def teardown_db(exception):
  #  mongo = g.pop('mongo', None)



app.register_blueprint(routes)


import db
