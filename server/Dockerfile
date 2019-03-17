FROM golang

ADD . /go/src/github.com/leepuppychow/heart_disease_prediction/server
WORKDIR /go/src/github.com/leepuppychow/heart_disease_prediction/server

RUN go install github.com/leepuppychow/heart_disease_prediction/server

ENTRYPOINT /go/bin/server

EXPOSE 8000