FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the built binary.
COPY ./.bin/app /root/app

# Copy all files under internal/infrastructure/configs
COPY ./internal/infrastructure/configs /root/internal/infrastructure/configs

# Copy all files under internal/infrastructure/drivers/postgres/setup
COPY ./internal/infrastructure/drivers/postgres/setup /root/internal/infrastructure/drivers/postgres/setup

CMD ["./app"]

# # Use the official Golang image to create a build artifact.
# FROM golang:1.18-alpine AS builder

# # Set necessary environment variables needed for the build
# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# # Move to working directory /build
# WORKDIR /build

# # Copy and download dependencies using go mod
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# # Copy the source code
# COPY . .

# # Build the application
# RUN go build -o app ./cmd/app

# # Use the official Alpine image for a lean production container.
# FROM alpine:latest

# # Install necessary packages
# RUN apk --no-cache add ca-certificates

# # Set working directory
# WORKDIR /root/

# # Copy the built binary from the builder stage
# COPY --from=builder /build/app /root/app

# # Copy the config directory
# COPY ./internal/infrastructure/configs /root/configs

# # Copy the .env file if needed for runtime configuration
# COPY .env /root/.env

# # Command to run the executable
# CMD ["./app"]
