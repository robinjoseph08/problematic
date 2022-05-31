FROM golang:1.18.2 AS problematic

WORKDIR /problematic

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go
COPY cmd cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /problematic/bin/problematic -installsuffix cgo -ldflags '-w -s' ./main.go
