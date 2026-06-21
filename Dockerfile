# =========================
# Builder
# =========================

FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates tzdata

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-s -w" \
    -o tinyurl \
    ./cmd/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-s -w" \
    -o migrate \
    ./cmd/migrate

# =========================
# Runtime
# =========================

FROM gcr.io/distroless/static-debian12

LABEL org.opencontainers.image.title="TinyURL Engineering Playground"
LABEL org.opencontainers.image.description="Production-inspired URL shortener built with Go, PostgreSQL, Redis, React and Docker"
LABEL org.opencontainers.image.source="https://github.com/abhinavkumar03/tinyurl-engineering-playground"
LABEL org.opencontainers.image.licenses="MIT"

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/tinyurl /tinyurl
COPY --from=builder /app/migrate /migrate

COPY --from=builder /app/internal/migration /internal/migration

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/tinyurl"]