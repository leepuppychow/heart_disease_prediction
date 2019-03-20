import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

def generate_histograms(data):
  df = pd.read_csv(data, sep=",")
  df.hist(bins=50, figsize=(20,15))
  plt.savefig("./images/histograms.jpg", format="jpg")

  # TODO: Maybe create separate histogram for each feature?
  # for column in df:
  #   feature = df[column]
  #   feature.hist(bins=50)
  #   filename = "./images/histogram-" + column + ".jpg"
  #   plt.autoscale(enable=True, axis="both")
  #   plt.savefig(filename, format="jpg")
  return