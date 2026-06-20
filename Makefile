APP_NAME=tinyurl

run:
	go run ./cmd/api

build:
	go build -o $(APP_NAME) ./cmd/api

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f $(APP_NAME)

docker-build:
	docker compose build

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

postgres:
	docker exec -it tinyurl-postgres psql -U postgres -d tinyurl

redis:
	docker exec -it tinyurl-redis redis-cli

migrate:
	sh scripts/migrate.sh

seed:
	sh scripts/seed.sh

check:
	go fmt ./...
	go vet ./...
	go test ./...

all:
	make check
	make build