#Load environment variables from .env file
include .env
export

postgres:
	 docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=kwabena@1 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb -U root -O root simple_bank

dropdb:
	docker exec -it postgres12 dropdb -U root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:kwabena@1@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateupps:
	migrate -path db/migration -database "${DATABASE_URL}" -verbose up

migratedownps:
	migrate -path db/migration -database "${DATABASE_URL}" -verbose down


migratedropps:
	migrate -path db/migration -database "${DATABASE_URL}" drop

migratedown:
	migrate -path db/migration -database "postgresql://root:kwabena@1@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres migrateup migratedown sqlc migrateupps migratedownps migrateupps1

