import numpy as np
import pandas as pd
from sklearn.preprocessing import MinMaxScaler
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.externals import joblib

def make_prediction(row):
  row = row.getvalue().split(",")
  row = list(map(lambda val: float(val), row))
  pickle = joblib.load('./models/heart_disease_logistic.pkl')
  model = pickle["model"]
  score = pickle["score"]
  prediction = int(model.predict([row])[0])

  return {
    "prediction": prediction,
    "score": score,
  }

def train_model(data):
  # TODO: Encode the categorical data!!
  df = pd.read_csv(data, sep=",")
  features = df[['age', 'sex', 'cp', 'trestbps', 'chol', 'fbs', 'restecg', 'thalach', 'exang', 'oldpeak', 'slope', 'ca', 'thal']]
  labels = df[['target']]

  scaler = MinMaxScaler()
  X = scaler.fit_transform(features)
  y = np.ravel(labels)

  X_train, X_test, y_train, y_test = train_test_split(X,y,test_size=0.3)
  logistic = LogisticRegression()
  logistic.fit(X_train,y_train)
  score = logistic.score(X_test,y_test)
  data = {
    "model": logistic,
    "score": score,
  }
  joblib.dump(data, "./models/heart_disease_logistic.pkl")

  return {
    'model_type': 'Logistic Regression Classifier',
    'score': score,
  }

