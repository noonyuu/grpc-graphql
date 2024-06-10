start:
	air -c .air.toml
test:
	go test ./...
build:
	go build -o cmd/api/run cmd/api/main.go
proto:
	protoc -I . ./user-service.proto --go_out=plugins=grpc:./pb
generate-models:
