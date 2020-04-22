BINARY_NAME=GoNgApp

build:
	go build -o bin/${BINARY_NAME}

test:
	go test -v ./...

run:
	build
	test
	./bin/${BINARY_NAME}

dev:
	go run server.go