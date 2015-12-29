FROM golang

RUN     apt-get update
RUN     apt-get install -y nodejs npm
RUN     ln -s /usr/bin/nodejs /usr/bin/node

ADD . /go

RUN     ./scripts/make_app.sh

RUN go install test-server/server
ENTRYPOINT /go/bin/server

EXPOSE 8080
