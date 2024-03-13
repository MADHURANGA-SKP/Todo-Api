postgres:
	docker run -d --name todoapi -p 5432:5432 -e POSTGRES_USER=pasan -e POSTGRES_PASSWORD=12345 postgres:16-alpine

createdb:
	docker exec -it todoapi createdb --username=pasan --owner=pasan todoc

dropdb:
	docker exec -it todoapi dropdb --username=pasan todoc

migrateup:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/todoc?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/todoc?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc