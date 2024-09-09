createdb:
	docker exec -it postgres1 createdb --username=nei --owner=nei simple_bank

dropdb:
	docker exec -it postgres1 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migratedown migrateup sqlc test