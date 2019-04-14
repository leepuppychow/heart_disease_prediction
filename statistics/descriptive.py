import pandas as pd

def get_stats(cache):
  stats = cache.get('stats')
  if stats is not None:
    return stats 
  return ""

def update_stats(data, cache):
  df = pd.read_csv(data, sep=",")
  data = df.describe().to_json(orient="columns")
  cache.set('stats', data) 
  return data
