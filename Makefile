all:
	go build && ./CommandLine
test:
	go test ./...
clean:
	rm -rf ./history ./CommandLine