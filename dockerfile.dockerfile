FROM golang:1.15.6

WORKDIR /go/src

COPY ./main ./

ENTRYPOINT ["./main"]