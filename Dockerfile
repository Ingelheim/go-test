FROM golang

RUN     apt-get update
RUN     apt-get install -y nodejs npm
RUN     ln -s /usr/bin/nodejs /usr/bin/node

ADD . /go

RUN     npm install -g bower
RUN     npm install -g grunt-cli
RUN     npm install
RUN     bower install --allow-root
RUN     grunt wiredep

RUN go install test-server/server
ENTRYPOINT /go/bin/server

EXPOSE 8080
