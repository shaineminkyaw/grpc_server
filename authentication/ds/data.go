package ds

import (
	"fmt"
	"grpc_basic/authentication/config"
	model "grpc_basic/authentication/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataSource struct {
	Sql *gorm.DB
}

var Auth_DB *gorm.DB

func AuthConnectToDB() *DataSource {
	conf := config.Init()
	host := conf.Host
	port := conf.Port
	dbname := conf.DB
	dbuser := conf.DBUser
	dbpassword := conf.DBPassword

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error on connecting to database!")
	} else {
		log.Printf("Connected Database :::")
	}

	Auth_DB = db
	db.AutoMigrate(
		&model.User{},
		&model.VerifyCode{},
		&model.UserToken{},
		&model.UserBankCard{},
		&model.UserCityTotalBankCard{},
	)

	return &DataSource{
		Sql: db,
	}
}
