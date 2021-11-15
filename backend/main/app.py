from flask import Flask
from flask_pymongo import PyMongo


app = Flask(__name__)
app.config["MONGO_URI"] = "mongodb://localhost:27017/test"
mongo = PyMongo(app)
basil = mongo.db.plants.find_one({"plantType": "Basil"})
print(basil)
@app.route("/")
def hello_world():
    return "Working"