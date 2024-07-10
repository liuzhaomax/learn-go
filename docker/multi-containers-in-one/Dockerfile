FROM golang:1.22.1 as builder
ADD . /workspace
WORKDIR /workspace

RUN env GOPROXY=https://goproxy.cn,direct \
    go build

FROM ubuntu:22.04

WORKDIR /workspace

COPY --from=builder /workspace/math-solver /workspace/