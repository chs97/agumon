FROM golang:latest as builder

WORKDIR /agumon-src

COPY . /agumon-src

RUN go mod download 

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -o /agumon

FROM alpine:latest

COPY --from=builder /agumon /

RUN  apk add g++

ENTRYPOINT ["/agumon"]