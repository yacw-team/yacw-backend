init:
	go mod download
run:
	go run main.go
build:
	GIN_MODE=release CGO_ENABLED=0 go build -o server -ldflags="-s -w" main.go