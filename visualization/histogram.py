import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

def generate_histogram(data):
  df = pd.read_csv(data, sep=",")
  print(df)
  return