all:
	go build ./...

init-dev-db:
	go run scripts/createdevdb.go config.yml

run-dev:
	go run httpserver/main.go config.yml