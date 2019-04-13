import pandas as pd

def stats(data):
  df = pd.read_csv(data, sep=",")
  return df.describe().to_json(orient="columns")
