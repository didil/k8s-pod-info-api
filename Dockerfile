# Base
FROM golang:1.19-alpine AS base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/k8s-pod-info-api .

# Final 
FROM alpine:3.17

WORKDIR /app

COPY --from=base /app/bin/k8s-pod-info-api /app/k8s-pod-info-api

CMD ["/app/k8s-pod-info-api"]