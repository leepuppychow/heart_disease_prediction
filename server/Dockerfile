FROM golang

ADD . /go/src/github.com/leepuppychow/heart_disease_prediction/server

RUN go get github.com/gomodule/redigo/redis
RUN go install github.com/leepuppychow/heart_disease_prediction/server

ENTRYPOINT /go/bin/server

EXPOSE 8000