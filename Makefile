proto:  
	rm -rf pb/*.go
	protoc --proto_path=./authentication/proto --go_out=./pb/ --go_opt=paths=source_relative \
    --go-grpc_out=./pb/ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./pb/  --grpc-gateway_opt=paths=source_relative \
    ./authentication/proto/*.proto
server:
	go run ./authentication/cmd/auth.go

evans:
	evans --host localhost --port 9088 -r repl

cert:
	./cert/gen.sh

.PHONY: proto server evans cert