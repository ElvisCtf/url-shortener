# Stage 1: Build
FROM golang:1.25.6-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o shorten-app .

# Stage 2: Runtime
FROM alpine:3.21
WORKDIR /root/

COPY --from=builder /app/shorten-app .

EXPOSE 8080

CMD ["./shorten-app"]
