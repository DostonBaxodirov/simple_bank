postgres:
	docker run --name postgres2 -p 5432:5432 -e POSTGRES_USER=nei -e POSTGRES_PASSWORD=54321 -d postgres:17.2-alpine3.21

createdb:
	docker exec -it postgres2 createdb --username=nei --owner=nei simple_bank

dropdb:
	docker exec -it postgres2 psql -U nei -c "DROP DATABASE simple_bank;"

migrateup:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://nei:54321@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go udemy-course/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
