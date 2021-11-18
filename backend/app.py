from flask import Flask
from main.routes import routes
from db import db

app = Flask(__name__)
app.config["MONGO_URI"] = "mongodb://localhost:27017/test"
db.get_connection(app)
app.register_blueprint(routes)


