postgresql:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine
createdb: 
	docker exec -it postgres14 createdb --username=root --owner=root commerce
dropdb: 
	docker exec -it postgres14 dropdb commerce
migrateup:
	migrate -path api/db/migration -database "postgresql://root:root@localhost/commerce?sslmode=disable" -verbose up
migratedown:
	migrate -path api/db/migration -database "postgresql://root:root@localhost/commerce?sslmode=disable" -verbose down
test:
	go test -v -cover ./...
	
.PHONY: postgresql createdb dropdb migrateup migratedown test
