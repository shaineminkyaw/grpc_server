package grpc

import (
	model "grpc_basic/authentication/model"
	"grpc_basic/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convert(user *model.User) *pb.User {
	//
	return &pb.User{
		Id:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		NationId:    user.NationID,
		City:        user.City,
		Bankcard:    user.BankCardNumber,
		Balance:     float32(user.Balance),
		Currency:    user.Currency,
		GenderType:  int32(user.Type),
		RegisterIp:  user.RegesiterIP,
		LastLoginIp: user.LastLoginIP,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		DeletedAt:   timestamppb.New(user.DeletedAt),
	}
}
