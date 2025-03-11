FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src

ADD . .

RUN go get -v
RUN go build -o /go/bin/tspl-print-api



FROM scratch

COPY --from=builder /go/bin/tspl-print-api /tspl-print-api

ENTRYPOINT ["/tspl-print-api"]
