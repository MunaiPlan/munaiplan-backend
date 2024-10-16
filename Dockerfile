FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest


# Copy the built binary.
COPY ./.bin/app /root/app

# Copy all files under internal/infrastructure/configs
COPY ./internal/infrastructure/configs /root/internal/infrastructure/configs

# Copy all files under internal/infrastructure/drivers/postgres/setup
COPY ./internal/infrastructure/drivers/postgres/setup /root/internal/infrastructure/drivers/postgres/setup

CMD ["./app"]
