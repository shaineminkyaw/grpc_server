package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"io"
	"io/ioutil"
	"log"

	"github.com/shaineminkyaw/grpc_server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//
func LoadTlsCredentials() (credentials.TransportCredentials, error) {
	//Load certificate of the CA who signed server certificate

	pemServerCA, err := ioutil.ReadFile("./cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, err
	}

	//Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil

}

func main() {
	//

	var opts []grpc.DialOption
	tlsCredentials, err := LoadTlsCredentials()
	if err != nil {
		log.Fatalf("cannot load tls credentials %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))

	conn, err := grpc.Dial("localhost:7070", opts...)
	if err != nil {
		log.Fatal("error on dial connection : ", err)
	}
	defer conn.Close()

	registerClient := pb.NewUserServiceClient(conn)

	fmt.Println("****** User Register Request and Response **********")

	ServerStream(registerClient)

}

func ServerStream(registerClient pb.UserServiceClient) {
	//
	var opts grpc.CallOption
	var registerDetails pb.RequestUserList = pb.RequestUserList{
		RequestData: make([]*pb.RequestUser, 0),
	}

	registerDetails.RequestData = append(registerDetails.RequestData, &pb.RequestUser{
		Email:      "smk20@gmail.com",
		Password:   "123456",
		VerifyCode: "7164",
		NationId:   "12/LMN(N)154655",
		GenderType: 2,
		City:       "mandalay",
	}, &pb.RequestUser{
		Email:      "smk21@gmail.com",
		Password:   "123456",
		VerifyCode: "6553",
		NationId:   "12/LMN(N)154655",
		GenderType: 1,
		City:       "yangon",
	})
	var i int64
	stream, err := registerClient.UserRegister(context.Background(), &registerDetails, opts)
	if err != nil {
		log.Fatalf("error on receive server data %v", err)
	}
	for {
		validUserList, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("recevie register stream error %v ", err)
		} else {
			i++
			log.Printf("%v valid User \n %v \n ", i, validUserList)
		}

	}
}
