migrate-up:
	migrate -path ./db/migration -database "postgres://root:root@localhost:5432/insert_service?sslmode=disable&application_name=insert_service" up

migrate-down:
	migrate -path ./db/migration -database "postgres://root:root@localhost:5432/insert_service?sslmode=disable&application_name=insert_service" down -all

run:
	go run main.go

.PHONY:
	migrate-up migrate-down