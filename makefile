buid:
	go run ./cmd/main.go
run:
	go run ./cmd/main.go
postgres:
	docker run --name psql -p 5432\:5432 -e POSTGRES_USER\=root -e POSTGRES_PASSWORD\=pass -d postgres\:13-alpine
psqlstart:
	docker start psql
psqlstop:
	docker stop psql
createdb:
	docker exec -it psql createdb --username=root --owner=root tmdb
dropdb:
	docker exec -it psql dropdb tmdb
migrateup:
	migrate -path migrations -database "postgresql://root:pass@localhost:5432/tmdb?sslmode=disable" -verbose up
migratedown:
	migrate -path migrations -database "postgresql://root:pass@localhost:5432/tmdb?sslmode=disable" -verbose down
.PHONY: run, buid

.DEFAULT_GOAL := buid