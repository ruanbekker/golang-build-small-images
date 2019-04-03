FROM golang:alpine

WORKDIR $GOPATH/src/mylekkepackage/mylekkeapp/
COPY app.go .
RUN go build -o /go/app

CMD ["/go/app"]
