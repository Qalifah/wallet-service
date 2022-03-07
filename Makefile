proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/wallet.proto

run-service:
	./run-service.sh

tests:
	go test ./handler