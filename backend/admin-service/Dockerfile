# Stage 1: Build
FROM golang:1.25-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app ./cmd/app

# Stage 2: Run
FROM alpine:3.21 AS final

COPY --from=builder /app /bin/app

EXPOSE 8082

CMD ["bin/app"]