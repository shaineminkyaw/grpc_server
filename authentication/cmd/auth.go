package main

import (
	"fmt"
	"grpc_basic/authentication/grpc"
)

func main() {
	fmt.Println("Hello World")
	go grpc.RunGatewayServer()
	grpc.RunGrpcServer()
}
