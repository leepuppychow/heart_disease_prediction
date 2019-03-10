FROM golang

ADD . /go/src/github.com/leepuppychow/heart_disease_prediction/prediction_old
WORKDIR /go/src/github.com/leepuppychow/heart_disease_prediction/prediction_old

RUN go get github.com/gomodule/redigo/redis
RUN go install github.com/leepuppychow/heart_disease_prediction/prediction_old

ENTRYPOINT /go/bin/prediction_old

EXPOSE 8080