RA512=4096
RA256=2048

.PHONY: update_local_deps
update_local_deps:
	go get -u github.com/swaggo/swag

# USAGE - make name=AddSomeDBTable db.migrate.create
db.migrate.create:
	goose -dir migrations create $(name) sql

db.migrate.up:
	goose -allow-missing -dir migrations postgres "postgres://postgres:postgres@localhost:5432/symbiosis?connect_timeout=180&sslmode=disable" up

db.migrate.down:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/symbiosis?connect_timeout=180&sslmode=disable" down

db.migrate.reset:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/symbiosis?connect_timeout=180&sslmode=disable" reset

.PHONY: swag
swag:
	swag init -g ./internal/api/api.go -o ./internal/api/swaggerdocs

create.keys:
	openssl genrsa -out keys/app.rsa $(RA512)
	openssl rsa -in keys/app.rsa -pubout -outform PEM -out keys/app.rsa.pub

generate:
	go generate ./...