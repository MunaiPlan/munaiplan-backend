# The build stage runs on the build host's architecture and cross-compiles for the target.
FROM --platform=$BUILDPLATFORM golang:1.21.1-alpine3.18 AS build
ARG TARGETOS=linux
ARG TARGETARCH
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags='-s -w' -o /out/app ./cmd/app

FROM alpine:3.23
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /out/app /app/app
COPY internal/infrastructure/configs/*.yml /app/internal/infrastructure/configs/
EXPOSE 8000
ENTRYPOINT ["/app/app"]
CMD ["serve"]
