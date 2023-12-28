FROM golang:alpine AS builder
ENV GO111MODULE=on
COPY . /api
WORKDIR /api
RUN go mod download
RUN go build 

FROM alpine AS api-runner
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
COPY --from=builder /api /api
WORKDIR /api
