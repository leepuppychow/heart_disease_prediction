from flask import Flask, jsonify, request
from flask_cors import CORS
import histogram
from io import StringIO

app = Flask(__name__)
CORS(app)

def convert_to_string(bytes_data):
  s = str(bytes_data, 'utf-8')
  return StringIO(s)

@app.route("/histogram", methods=['POST'])
def histogram_handler():
  histogram.generate_histogram(convert_to_string(request.data))
  return "HISTOGRAM HANDLER - PYTHON"

if __name__ == "__main__":
  app.run(host="0.0.0.0", port=8888)