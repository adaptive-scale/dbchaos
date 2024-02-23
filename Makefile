BINARY_NAME=dbchaos
setup:
	go mod tidy
	rm -rf bin
	mkdir -p bin
build: setup
	go build -o bin/$(BINARY_NAME) .