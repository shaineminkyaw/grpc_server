package config

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/ini.v1"
)

type app struct {
	AppAddress string
}

type mysql struct {
	DBPort     string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
}

type grpc struct {
	AuthGrpc string
}

var (
	App               app
	SQL               mysql
	PrivateKey        *rsa.PrivateKey
	PublicKey         *rsa.PublicKey
	GRPC              grpc
	AuthSecretKey     string
	ServerCredentials grpc_pem
	ServerCert        *rsa.PrivateKey
	ServerKey         *rsa.PublicKey
)

type grpc_pem struct {
	ServerCert string
	ServerKey  string
}

type Config struct {
	AppAddress  string
	GRPCAddress string
	Port        string
	Host        string
	DB          string
	DBUser      string
	DBPassword  string
	Private     *rsa.PrivateKey
	Public      *rsa.PublicKey
	SecretKey   string
	ServerCert  *rsa.PrivateKey
	ServerKey   *rsa.PublicKey
}

func Init() *Config {
	auth_iniPath := "authentication/config/config.ini"

	args := os.Args
	if len(args) > 1 {
		auth_iniPath = args[1]

	}

	auth_iniFile, err := ini.Load(auth_iniPath)
	if err != nil {
		log.Fatalf("Load %v error %v \n", auth_iniFile, err.Error())
		os.Exit(1)
	}

	// app
	auth_app := auth_iniFile.Section("auth_app")
	App.AppAddress = auth_app.Key("Auth_HTTPAddress").String()

	//Server Credentials
	// serverCred := auth_iniFile.Section("grpc_credentials")
	// ServerCredentials.ServerCert = serverCred.Key("Server_Cert").String()
	// cert, err := ioutil.ReadFile(ServerCredentials.ServerCert)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// certKey, err := jwt.ParseRSAPrivateKeyFromPEM(cert)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// ServerCert = certKey
	// ServerCredentials.ServerKey = serverCred.Key("Server_Key").String()
	// key, err := ioutil.ReadFile(ServerCredentials.ServerKey)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// encrptKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// ServerKey = encrptKey

	//grpc
	auth_grpc := auth_iniFile.Section("auth_grpc")
	GRPC.AuthGrpc = auth_grpc.Key("Auth_GRPC").String()

	//mysql
	auth_mysql := auth_iniFile.Section("auth_mysql")
	SQL.DBHost = auth_mysql.Key("Auth_Host").String()
	SQL.DBPort = auth_mysql.Key("Auth_Port").String()
	SQL.DBName = auth_mysql.Key("Auth_DBName").String()
	SQL.DBUser = auth_mysql.Key("Auth_Username").String()
	SQL.DBPassword = auth_mysql.Key("Auth_Password").String()

	//rsa
	auth_rsa := auth_iniFile.Section("auth_rsa")
	//secret key
	AuthSecretKey = auth_rsa.Key("Auth_Secret_Key").String()

	//private key
	prvKey := auth_rsa.Key("Auth_Private_KeyFile").String()
	prvByte, err := ioutil.ReadFile(prvKey)
	if err != nil {
		log.Fatalf("error on read file %v", err.Error())
	}

	private, err := jwt.ParseRSAPrivateKeyFromPEM(prvByte)
	if err != nil {
		log.Fatalf("error on change pem file to rsa private key %v", err.Error())
	}
	//rsa private key
	PrivateKey = private

	pubKey := auth_rsa.Key("Auth_Public_KeyFile").String()
	pubByte, err := ioutil.ReadFile(pubKey)
	if err != nil {
		log.Fatalf("error on read file %v", err.Error())
	}

	public, err := jwt.ParseRSAPublicKeyFromPEM(pubByte)
	if err != nil {
		log.Fatalf("error on change pem to rsa public key  %v ", err.Error())
	}

	//rsa public key
	PublicKey = public

	return &Config{
		AppAddress:  App.AppAddress,
		GRPCAddress: GRPC.AuthGrpc,
		Port:        SQL.DBPort,
		Host:        SQL.DBHost,
		DB:          SQL.DBName,
		DBUser:      SQL.DBUser,
		DBPassword:  SQL.DBPassword,
		Private:     PrivateKey,
		Public:      PublicKey,
		SecretKey:   AuthSecretKey,
	}

}
