from flask import Flask, jsonify, request
from flask_cors import CORS
from io import StringIO
import descriptive

app = Flask(__name__)
CORS(app)

def convert_to_string(bytes_data):
  s = str(bytes_data, 'utf-8')
  return StringIO(s)

@app.route("/stats", methods=['POST'])
def descriptive_stats():
  resp = descriptive.stats(convert_to_string(request.data))
  return resp, 200

if __name__ == "__main__":
  app.run(host="0.0.0.0", port=8111)