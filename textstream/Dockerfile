# Docker image used for building and running unit tests for the Flink Operator.
#
# It installs required build dependencies (e.g., Go 12+), copies the project
# source code into the container, build and run tests.
#
# Usage: 
#
# docker build -t flink-operator-builder -f Dockerfile.builder .
# docker run flink-operator-builder


FROM ubuntu:18.04

RUN apt update && apt install -yqq curl git make gcc

# Install Go.
RUN curl -s https://dl.google.com/go/go1.12.10.linux-amd64.tar.gz | tar -xz -C /usr/local/
ENV GOROOT=/usr/local/go
ENV PATH=${PATH}:${GOROOT}/bin

# Download source code.
RUN git clone https://github.com/functicons/go-examples.git

WORKDIR /go-examples/textstream

# Build binary.
RUN go build -a -o ./bin/textstream text_stream.go

EXPOSE 8080
ENTRYPOINT [ "./bin/textstream" ]
