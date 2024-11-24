all:
	go build && ./CommandLine
test:
	go test ./...
clear:
	rm -rf ./history ./CommandLine