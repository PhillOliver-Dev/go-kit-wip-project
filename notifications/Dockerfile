FROM golang

RUN mkdir -p /go/src/kit-test

ADD . /go/src/kit-test
WORKDIR /go/src/kit-test

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run kit-test/notifications/cmd  -watch kit-test/notifications
