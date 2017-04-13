FROM golang:1.8-alpine
ARG name
ENV CGO_ENABLED 0
ENV GOOS linux
EXPOSE 8080

RUN apk add --no-cache --update make git mercurial

RUN mkdir -p $GOPATH/src/github.com/alexbt/go-consul-example/${name}
ADD . $GOPATH/src/github.com/alexbt/go-consul-example/${name}
WORKDIR $GOPATH/src/github.com/alexbt/go-consul-example/${name}

RUN make

RUN rm -fr $GOPATH/src
RUN ls -lrt $GOPATH/bin/

ENTRYPOINT $GOPATH/bin/service-one