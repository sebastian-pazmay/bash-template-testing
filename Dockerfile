ARG OS=ubuntu:18.04

FROM golang:1.14.15-alpine3.12 AS build
WORKDIR /go/src/
COPY bash-testing-app/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/bashTest main.go

FROM ${OS}
WORKDIR /root/
COPY --from=build /tmp/bashTest .
CMD ["./bashTest"]  
