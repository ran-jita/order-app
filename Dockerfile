# syntax=docker/dockerfile:1.0-experimental
FROM golang:1.16-alpine as builder
ARG APP_ENV
# Set workdir
WORKDIR /order-app
# Copy go mod and go sum
COPY go.mod .
COPY go.sum .
# Get dependencies
RUN go mod download
# Add files from repo
ADD . .
# Update and install git
RUN apk update --no-cache && \
    apk add --no-cache git && \
    # Create folder
    mkdir -p /src/app && \
    # Build go application
    export GO111MODULE=on && \
    cd app && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /src/order-app

# Scratch
FROM alpine:3.12
# Set workdir
WORKDIR /src
# Copy from builder
COPY --from=builder /src/order-app /src/order-app
COPY --from=builder /src/.env /src/.env

# Install tzdata and update alpine
RUN apk update --no-cache && \
    apk add --no-cache bash && \
    apk add --no-cache -U tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    # Set permission
    chmod +x /src/order-app && \
    # remove unused files
    rm -rf Dockerfile \
    dockerfile \
    .git \
    .env.* \
    deployment

# Expose application
EXPOSE 8000
# Run application
CMD ["/src/order-app"]