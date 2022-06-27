postgres:
		docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
		docker exec -it postgres14 createdb --username=root --owner=root mini_bank

dropdb:
	docker exec -it postgres14 dropdb mini_bank

migrationup:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/mini_bank?sslmode=disable" -verbose up
		   
migrationdown:
			migrate -path db/migration -database "postgresql://root:secret@localhost:5432/mini_bank?sslmode=disable" -verbose down
sqlc:
	 sqlc generate

test:
	 go test -v -cover ./...	

server:
	 go run main.go	

.phony:postgres createdb dropdb migrationup migrationdown sqlc test server