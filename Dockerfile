from golang:1.10

ADD . /go/src/github.com/kodek/hello-travis-docker

RUN go install github.com/kodek/hello-travis-docker/cmd

CMD ["main"]
