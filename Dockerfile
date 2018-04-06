FROM golang:1.10-alpine

ENV GOPATH /go
ENV USER root

RUN apk update && apk add git make curl

# pre-install known dependencies before the source, so we don't redownload them whenever the source changes
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
 && govendor get github.com/hawky-4s-/camunda-rest-client-go

COPY . $GOPATH/src/github.com/hawky-4s-/camunda-rest-client-go

RUN cd $GOPATH/src/github.com/hawky-4s-/camunda-rest-client-go \
 	&& make install test
