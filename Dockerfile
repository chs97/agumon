FROM golang:latest as builder

WORKDIR /agumon-src

COPY . /agumon-src

RUN go mod download 

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -o /agumon

FROM golang:alpine

COPY --from=builder /agumon /

RUN  apk add g++ gcc openjdk8
ENV JAVA_HOME=/usr/lib/jvm/java-1.8-openjdk
ENV PATH="$JAVA_HOME/bin:${PATH}" 

ENTRYPOINT ["/agumon"]