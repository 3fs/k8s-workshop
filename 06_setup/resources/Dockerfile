FROM golang:1.12.4-stretch AS builder

WORKDIR /app

COPY ./code .

RUN go build -o example-server main.go

# =============================================================================
FROM gcr.io/distroless/base:latest AS deployable

COPY --from=builder /app/example-server /usr/bin/example-server

# Sourced from https://github.com/arcadeJHS/HTML5-Space-Invaders
COPY resources/game/index.html /game/index.html

ENTRYPOINT ["/usr/bin/example-server"]
