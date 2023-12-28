FROM golang:alpine AS builder
ENV GO111MODULE=on
COPY . /api
WORKDIR /api
RUN go mod download
RUN go build 

FROM alpine AS api-runner
COPY --from=builder /api /api
WORKDIR /api
