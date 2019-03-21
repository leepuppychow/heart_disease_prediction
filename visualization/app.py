from flask import Flask, jsonify, request, send_from_directory, make_response, jsonify
from flask_cors import CORS
import histogram
from io import StringIO

app = Flask(__name__, static_url_path='')
CORS(app)

def convert_to_string(bytes_data):
  s = str(bytes_data, 'utf-8')
  return StringIO(s)

@app.route("/histograms", methods=['POST'])
def create_histograms():
  histogram.generate_histograms(convert_to_string(request.data))
  return make_response(jsonify( {'message': 'Successfully generated histograms'} ), 201)

@app.route("/histograms", methods=['GET'])
def all_histograms():
  return send_from_directory("images", "histograms.jpg")
  

@app.route("/histograms/<feature>", methods=['GET'])
def show_histogram(feature):
  return send_from_directory("images", "histogram-" + feature + ".jpg")

if __name__ == "__main__":
  app.run(host="0.0.0.0", port=8888)

# Eventually could just save images in S3?