import pandas as pd

def update_stats(data):
  df = pd.read_csv(data, sep=",")
  descriptive = df.describe().to_json(orient="columns")
  correlations = df.corr().to_json(orient="columns")
  return (descriptive, correlations)
