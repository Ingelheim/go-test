FROM golang

ADD . /go
RUN go install test-server/server
ENTRYPOINT /go/bin/server

EXPOSE 8080
