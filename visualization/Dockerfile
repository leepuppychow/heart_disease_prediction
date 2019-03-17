FROM golang

ADD . /go/src/github.com/leepuppychow/heart_disease_prediction/visualization
WORKDIR /go/src/github.com/leepuppychow/heart_disease_prediction/visualization

RUN go get github.com/ajstarks/svgo
RUN go install github.com/leepuppychow/heart_disease_prediction/visualization

ENTRYPOINT /go/bin/visualization

EXPOSE 8888