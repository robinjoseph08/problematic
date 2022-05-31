FROM golang:1.18.2 AS problematic

WORKDIR /problematic
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /problematic/bin/problematic -installsuffix cgo -ldflags '-w -s' ./main.go
