build: build-swaager
	go build -o bin/ecom

build-swaager:
	swag init -g router/router.go

test:
	go test -v -cover ./...