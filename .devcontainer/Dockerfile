FROM golang:1-bullseye
WORKDIR /workspace
COPY go.mod .
COPY go.sum .
RUN go mod download
