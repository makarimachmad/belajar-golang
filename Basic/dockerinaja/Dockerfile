FROM golang:latest

#MAINTAINER achmad.widyanto@spesolution.com
ADD . /go/src/belajar-golang/basic/dockerinaja

ARG appname=e-document-api
ARG http_port=1323

RUN go get -d github.com/go-sql-driver/mysql
RUN go get -d github.com/labstack/echo
RUN . go install dockerinaja@latest

ENTRYPOINT /go/bin/belajar-golang/Basic/dockerinaja

ENV PORT $http_port
ENV DB_HOST localhost
ENV DB_PORT 3306
ENV DB_USER root
ENV DB_PASS password

EXPOSE $http_port

RUN mkdir -p /go/src/belajar-golang/Basic/dockerinaja
COPY . /go/src/belajar-golang/Basic/dockerinaja
WORKDIR /go/src/belajar-golang/Basic/dockerinaja

CMD go run main.go