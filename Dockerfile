FROM golang:alpine

WORKDIR /src
COPY src ./

ENV CGO_ENABLED 0
ENV GOOS linux

RUN go mod download
ENTRYPOINT ["go", "test", "-v"]
