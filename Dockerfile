FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the built binary.
COPY ./.bin/app /root/app

# Copy the catalog files directory
COPY ./internal/catalog/catalog_files/ /root/catalog_files/

CMD ["./app"]