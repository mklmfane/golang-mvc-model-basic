FROM golang:1.24.4

WORKDIR /app

COPY go.mod go.sum ./

RUN apt-get update && apt-get install -y ca-certificates
RUN go env -w GOPROXY=direct
RUN go mod download

COPY . .
