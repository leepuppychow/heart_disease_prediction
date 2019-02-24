FROM golang

ADD . /go/src/github.com/leepuppychow/heart_disease_prediction/prediction_service

WORKDIR /go/src/github.com/leepuppychow/heart_disease_prediction/prediction_service

RUN go get github.com/gomodule/redigo/redis
RUN go install github.com/leepuppychow/heart_disease_prediction/prediction_service

ENTRYPOINT /go/bin/prediction_service