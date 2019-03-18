import numpy as np
import pandas as pd
from sklearn.preprocessing import MinMaxScaler
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
# joblib is for model persistence so it doesn't have to train model at every request
from sklearn.externals import joblib

def make_prediction(data):
  df = pd.read_csv(data, sep=",")
  print(df)
  return

def train_model(data):
  df = pd.read_csv(data, sep=",")
  print(df)
  return

# EXAMPLE FROM MY FALL RISK API
# def train():
  # df = pd.read_csv("some_file.csv")
  # features = df[['age','berg_balance','gait_speed']]
  # labels = df[['fall_risk']]

  # # Perform feature scaling
  # scaler = MinMaxScaler()
  # X = scaler.fit_transform(features)
  # y = np.ravel(labels)

  # # 70:30 training:testing data split
  # X_train, X_test, y_train, y_test = train_test_split(X,y,test_size=0.3)
  # logistic = LogisticRegression()
  # logistic.fit(X_train,y_train)
  # score = logistic.score(X_test,y_test)

  # data = {
  #     "model": logistic,
  #     "score": score,
  #     "age_min": df[['age']].min(),
  #     "berg_min": df[['berg_balance']].min(),
  #     "gait_min": df[['gait_speed']].min(),
  #     "age_range": df[['age']].max() - df[['age']].min(),
  #     "berg_range": df[['berg_balance']].max() - df[['berg_balance']].min(),
  #     "gait_range": df[['gait_speed']].max() - df[['gait_speed']].min()
  # }

  # # Pickle the model, do you don't train everytime API is called
  # joblib.dump(data, "logistic_regression_data.pkl")

# def predict(age, berg, gait):
#     data = joblib.load("./logistic_regression_data.pkl")

#     age_scaled = ((age - data["age_min"]) / data["age_range"]).item()
#     berg_scaled = ((berg - data["berg_min"]) / data["berg_range"]).item()
#     gait_scaled = ((gait - data["gait_min"]) / data["gait_range"]).item()

#     prediction = data["model"].predict([[age_scaled, berg_scaled, gait_scaled]])

#     if prediction == 0:
#       return {"fall_risk": "LOW",
#         "model_accuracy": data["score"]}
#     else:
#       return {"fall_risk": "HIGH",
#         "model_accuracy": data["score"]}