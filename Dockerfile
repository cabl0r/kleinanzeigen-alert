FROM golang:alpine
RUN apk update && apk add --no-cache git gcc musl-dev
ADD . /go/src/github.com/cabl0r/kleinanzeigen-alert/
WORKDIR /go/src/github.com/cabl0r/kleinanzeigen-alert/
RUN go get -d -v
RUN go install github.com/cabl0r/kleinanzeigen-alert/

ENTRYPOINT ["/go/bin/kleinanzeigen-alert"]