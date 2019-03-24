from flask import Flask, jsonify, request
from flask_cors import CORS
import logistic
from io import StringIO

app = Flask(__name__)
CORS(app)

def convert_to_string(bytes_data):
  s = str(bytes_data, 'utf-8')
  return StringIO(s)

@app.route("/train", methods=['POST'])
def train():
  resp = logistic.train_model(convert_to_string(request.data))
  return jsonify(resp), 200

@app.route("/predict", methods=['POST'])
def predict():
  resp = logistic.make_prediction(convert_to_string(request.data))
  return jsonify(resp), 200

if __name__ == "__main__":
  app.run(host="0.0.0.0", port=8080)