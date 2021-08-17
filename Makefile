build: build-swaager
	go build -o bin/ecom

build-swaager:
	swag init --parseDependency --parseInternal

test:
	go test -v -cover ./...