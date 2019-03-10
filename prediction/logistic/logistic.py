import numpy as np
import pandas as pd
from sklearn.preprocessing import MinMaxScaler
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
# joblib is for model persistence so it doesn't have to train model at every request
from sklearn.externals import joblib

def logistic():
  # EXAMPLE FROM MY FALL RISK API

  # df = pd.read_csv('fall_risk.csv')
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
  return ""