ARG OS=ubuntu:18.04

FROM golang:1.14.15-alpine3.12 AS build
WORKDIR /go/src/
COPY os-testing-app/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/osTest main.go

FROM ${OS}
WORKDIR /root/
COPY --from=build /tmp/osTest .
CMD ["./osTest"]  
