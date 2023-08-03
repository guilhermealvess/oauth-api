run:
	go run cmd/main.go

deps:
	go mod tidy

queries-generate:
	sqlc generate

protoc:
	protoc --go_out=. --go-grpc_out=. proto/api_server.proto
