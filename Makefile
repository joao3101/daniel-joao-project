docker_postgres:
	docker pull postgres:12-alpine

docker_compose:
	docker-compose up -d

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it fantasy createdb --username=root --owner=root fantasy

dropdb:
	docker exec -it fantasy dropdb fantasy

sqlboiler:
	sqlboiler psql 

#lembrar de colocar o nome após o comando
migrate_create:
	migrate create -ext sql -dir db/migrations -seq

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5431/fantasy?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5431/fantasy?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5431/fantasy?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5431/fantasy?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: docker_postgres docker_compose sqlboiler createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock