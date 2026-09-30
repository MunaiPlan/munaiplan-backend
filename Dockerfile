FROM golang:1.21.1-alpine3.18 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags='-s -w' -o /out/app ./cmd/app

FROM alpine:3.23
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /out/app /app/app
COPY internal/infrastructure/configs/main.yml /app/internal/infrastructure/configs/main.yml
EXPOSE 8000
ENTRYPOINT ["/app/app"]
CMD ["serve"]
