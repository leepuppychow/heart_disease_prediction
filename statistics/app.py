from flask import Flask, jsonify, request
from flask_caching import Cache
from flask_cors import CORS
from io import StringIO
import descriptive

config = {
    "DEBUG": True,
    "CACHE_TYPE": "simple",
    "CACHE_DEFAULT_TIMEOUT": 300
}

app = Flask(__name__)
CORS(app)

app.config.from_mapping(config)
cache = Cache(app)

def convert_to_string(bytes_data):
  s = str(bytes_data, 'utf-8')
  return StringIO(s)

@app.route("/stats", methods=['GET'])
def get_current_stats():
  return descriptive.get_stats(cache), 200

@app.route("/stats", methods=['POST'])
def create_new_descriptive_stats():
  data = convert_to_string(request.data)
  resp = descriptive.update_stats(data, cache)
  return resp, 200

if __name__ == "__main__":
  app.run(host="0.0.0.0", port=8111)