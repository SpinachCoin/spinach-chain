build:
	go build -o bin/spinach-chain ./src

start:
	./bin/spinach-chain

test:
	go test ./tests/...