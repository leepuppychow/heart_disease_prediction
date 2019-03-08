FROM golang

ADD . /go/src/github.com/leepuppychow/heart_disease_prediction/prediction
WORKDIR /go/src/github.com/leepuppychow/heart_disease_prediction/prediction

RUN go get github.com/gomodule/redigo/redis
RUN go install github.com/leepuppychow/heart_disease_prediction/prediction

ENTRYPOINT /go/bin/prediction

EXPOSE 8080