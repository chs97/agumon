FROM golang:latest as builder

WORKDIR /agumon

COPY . /agumon

RUN go mod download && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -o /agumon

FROM alpine:latest

COPY --from=builder /agumon /

ENTRYPOINT ["/agumon"]