FROM golang:alpine

WORKDIR /src
COPY . ./

ENV CGO_ENABLED 0
ENV GOOS linux
