run-api:
	go run ./cmd/api

swagger:
	GOTOOLCHAIN=go1.23.5 go run github.com/swaggo/swag/cmd/swag@v1.16.6 init -g cmd/api/main.go -o docs

test:
	go test ./... -count=1 -race
	golangci-lint run -v ./...
