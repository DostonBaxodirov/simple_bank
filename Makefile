createdb:
	docker exec -it postgres1 createdb --username=nei --owner=nei simple_bank

dropdb:
	docker exec -it postgres1 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simpleBank/tutorial Store

.PHONY: createdb dropdb migratedown migrateup sqlc test server mock migratedown1 migrateup1
