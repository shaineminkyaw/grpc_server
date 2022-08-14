proto:  
	rm -rf pb/*.go
	protoc --proto_path= github.com/shaineminkyaw/grpc_server/authentication/proto --go_out=github.com/shaineminkyaw/grpc_server/pb/ --go_opt=paths=source_relative \
    --go-grpc_out=github.com/shaineminkyaw/grpc_server/pb/ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=github.com/shaineminkyaw/grpc_server/pb/  --grpc-gateway_opt=paths=source_relative \
    github.com/shaineminkyaw/grpc_server/authentication/proto/*.proto
server:
	go run ./authentication/cmd/auth.go

evans:
	evans --host localhost --port 9088 -r repl

cert:
	./cert/gen.sh

.PHONY: proto server evans cert