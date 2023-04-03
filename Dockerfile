FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git tzdata
RUN apk add ca-certificates && update-ca-certificates

WORKDIR $GOPATH/src/nordend/

COPY go.mod .

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /go/bin/server cmd/server/*.go


FROM ubuntu:latest

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/server /go/bin/server

WORKDIR /go/bin/

COPY ./entrypoint.sh .

RUN apt-get update
RUN apt-get install -y jq

RUN chmod +x entrypoint.sh