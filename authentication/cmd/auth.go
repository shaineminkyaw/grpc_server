package main

import (
	"fmt"

	"github.com/shaineminkyaw/grpc_server/authentication/grpc"
)

func main() {
	fmt.Println("Hello World")
	go grpc.RunGatewayServer()
	grpc.RunGrpcServer()
}
