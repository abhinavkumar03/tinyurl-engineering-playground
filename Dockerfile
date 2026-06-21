FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
-o tinyurl \
./cmd/api

RUN CGO_ENABLED=0 GOOS=linux go build \
-o migrate \
./cmd/migrate

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/tinyurl .

COPY --from=builder /app/migrate .

EXPOSE 8080

CMD ["./tinyurl"]