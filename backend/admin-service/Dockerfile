# Stage 1: Build
FROM golang:1.25-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# 1. Copy go.work and all module definitions to leverage Docker caching
COPY go.work ./
COPY shared/go.mod shared/go.sum ./shared/
COPY admin-service/go.mod admin-service/go.sum ./admin-service/
COPY redirect-service/go.mod redirect-service/go.sum ./redirect-service/
COPY shorten-service/go.mod shorten-service/go.sum ./shorten-service/

# 2. Download dependencies for the specific service module
WORKDIR /build/admin-service
RUN go mod download

# 3. Copy the actual source code for both modules
WORKDIR /build
COPY shared/ ./shared/
COPY admin-service/ ./admin-service/

# 4. Build from the service directory
WORKDIR /build/admin-service
RUN go build -o /app ./cmd/app/main.go

# Stage 2: Run
FROM alpine:3.21 AS final

COPY --from=builder /app /bin/app

EXPOSE 8082

CMD ["/bin/app"] 
