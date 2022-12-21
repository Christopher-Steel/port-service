FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o port-service cmd/port-service/main.go
FROM alpine
WORKDIR /build
COPY --from=builder /build/port-service /build/port-service
CMD ["./port-service"]