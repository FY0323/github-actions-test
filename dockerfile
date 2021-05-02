FROM golang:1.15.6 as builder

WORKDIR /go/src

COPY ./main.go ./
RUN go build -o /go/bin/main


FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

ENV PORT 8080

CMD ["/app/main"]