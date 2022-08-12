package grpc

import (
	"context"
	"fmt"

	"math/rand"
	"time"

	model "github.com/shaineminkyaw/grpc_server/authentication/model"
	"github.com/shaineminkyaw/grpc_server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (server *AuthServer) GetVerifyCode(ctx context.Context, req *pb.RequestVerifyCode) (*pb.ResponseVerifyCode, error) {

	//
	number := time.Now().UnixNano()
	rand.Seed(number)
	code := fmt.Sprintf("%v%v%v%v", rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))
	expTime := time.Now().Add(time.Minute * 30)
	mail := req.Email
	cUser := &model.VerifyCode{
		Email:      mail,
		Code:       code,
		ExpireTime: expTime,
	}
	verfiy := &model.VerifyCode{}
	err := server.Database.Sql.Model(&model.VerifyCode{}).Where("email = ?", req.GetEmail()).First(&verfiy).Error
	if err == gorm.ErrRecordNotFound {
		err := server.Database.Sql.Model(&model.VerifyCode{}).Create(&cUser).Error
		// log.Printf("create data to table ....%v", cUser)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error on verify code create : %v", err.Error())
		}
	} else {
		data := server.Database.Sql.Model(&model.VerifyCode{}).Where("email =?", req.GetEmail())
		err = data.Updates(&cUser).Error
		// log.Printf("updates data to table ....%v", cUser)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error on verify code updates %v ", err.Error())
		}
	}

	return &pb.ResponseVerifyCode{
		Code: cUser.Code,
	}, nil

	//

	// return nil, status.Errorf(codes.Unimplemented, "method GetVerifyCode not implemented")
}
