.PHONY: migrate-up
migrate-up:
  docker compose -f docker-compose.yml --profile tools run --rm migrate up

.PHONY: migrate-down
migrate-down:
  docker compose -f docker-compose.yml --profile tools run --rm migrate down 1

.PHONY: migrate-create
migrate-create:
  docker compose -f docker-compose.yml --profile tools run --rm migrate create  -ext sql -dir /migrations create_persons_table

