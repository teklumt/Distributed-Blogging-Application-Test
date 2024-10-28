# Build the API Gateway
FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o api-gateway main.go

# Run the API Gateway
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api-gateway .
COPY .env .env
EXPOSE 8080
CMD ["./api-gateway"]
