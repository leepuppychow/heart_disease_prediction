import pandas as pd

def update_stats(data):
  df = pd.read_csv(data, sep=",")
  data = df.describe().to_json(orient="columns")
  return data
