from flask import Flask, jsonify, request
from redis import Redis, RedisError
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

@app.route("/train")
def train():
  response = "HELLO TRAIN HANDLER - PYTHON"
  print(response)
  return response

@app.route("/predict")
def predict():
  response = "HELLO PREDICT HANDLER - PYTHON"
  print(response)
  return response

if __name__ == "__main__":
  app.run(host="0.0.0.0", port=8080)