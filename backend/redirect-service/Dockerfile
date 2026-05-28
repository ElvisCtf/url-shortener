# Stage 1: Build
FROM golang:1.25-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# 1. Copy go.work and all module definitions to leverage Docker caching
COPY go.work ./
COPY shared/go.mod shared/go.sum ./shared/
COPY redirect-service/go.mod redirect-service/go.sum ./redirect-service/
COPY shorten-service/go.mod shorten-service/go.sum ./shorten-service/
COPY admin-service/go.mod admin-service/go.sum ./admin-service/

# 2. Download dependencies for the specific service module
WORKDIR /build/redirect-service
RUN go mod download

# 3. Copy the actual source code for both modules
WORKDIR /build
COPY shared/ ./shared/
COPY redirect-service/ ./redirect-service/

# 4. Build from the service directory
WORKDIR /build/redirect-service
RUN go build -o /app ./main.go

# Stage 2: Run
FROM alpine:3.21 AS final

COPY --from=builder /app /bin/app

EXPOSE 8081

CMD ["/bin/app"] 
