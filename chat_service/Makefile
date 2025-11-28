run:
	go run cmd/api/main.go

db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dsn) goose -dir=$(migrationPath) status

swag:
	cd cmd/api && \
	swag init