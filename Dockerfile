# Build stage
FROM golang:1.25.0-alpine AS builder

WORKDIR /app
COPY . .

# Install dependencies and build
RUN go mod download
RUN go build -o /chip-api

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /chip-api /app/chip-api

EXPOSE 8080
CMD ["/app/chip-api"]