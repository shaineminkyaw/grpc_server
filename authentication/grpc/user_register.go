package grpc

import (
	"context"
	"log"
	"time"

	"github.com/mazen160/go-random"
	model "github.com/shaineminkyaw/grpc_server/authentication/model"
	"github.com/shaineminkyaw/grpc_server/authentication/util"
	"github.com/shaineminkyaw/grpc_server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (server *AuthServer) UserRegister(ctx context.Context, req *pb.RequestUser) (*pb.ResponseUser, error) {
	//

	vUser := &model.VerifyCode{}
	//@@@validate verifycode
	err := server.Database.Sql.Model(&model.VerifyCode{}).Where("email = ?", req.GetEmail()).First(&vUser).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error on validate verify code ")
	}
	if vUser == nil {
		return nil, status.Errorf(codes.NotFound, "email not registered for verify code !")
	}
	if req.GetVerifyCode() != vUser.Code || time.Now().Unix() > vUser.ExpireTime.UnixNano() {
		return nil, status.Errorf(codes.Internal, "verifycode invalid")
	}

	//hash password
	userPassword, err := util.HashPassword(req.GetPassword())
	log.Println("User Password ....", userPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error on hash password")
	}
	charset := "12345678"
	length := 8
	usr, _ := random.Random(length, charset, true)
	userName := "U" + usr
	currency := "USD"
	bankCard, err := util.GetBankCardNumber(req.City)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erorr on generate bank card")
	}
	user := &model.User{
		Username:       userName,
		Password:       userPassword,
		Email:          req.GetEmail(),
		RegesiterIP:    "",
		LastLoginIP:    "",
		NationID:       req.GetNationId(),
		BankCardNumber: bankCard,
		City:           req.GetCity(),
		Balance:        0,
		Currency:       currency,
		Type:           int8(req.GetGenderType()),
	}

	fUser := &model.User{}
	data := server.Database.Sql.Model(&model.User{}).Where("email =?", req.GetEmail())
	err = data.First(&fUser).Error
	if err == gorm.ErrRecordNotFound {
		err := server.Database.Sql.Model(&model.User{}).Create(&user).Error
		if err != nil {
			return nil, status.Errorf(codes.Internal, "erorr on user create !")
		}
		err = util.SaveUserBankCard(req.GetCity(), user.ID, bankCard)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "erorr on store bank card ")
		}
	} else {
		return nil, status.Errorf(codes.Internal, "user already exists")
	}

	resp := &pb.ResponseUser{
		User:      convert(user),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}

	return resp, nil
}

// func (server *AuthServer) UserRegister(req *pb.RequestUserList, stream pb.UserService_UserRegisterServer) error {
// 	respUser := make([]*pb.ResponseUser, 0)
// 	for _, userData := range req.RequestData {

// 		// registerUser := &pb.RequestUser{
// 		// 	Email:      userData.Email,
// 		// 	Password:   userData.Password,
// 		// 	VerifyCode: userData.VerifyCode,
// 		// 	NationId:   userData.NationId,
// 		// 	GenderType: userData.GenderType,
// 		// 	City:       userData.City,
// 		// }

// 		vUser := &model.VerifyCode{}
// 		//@@@validate verifycode
// 		err := server.Database.Sql.Model(&model.VerifyCode{}).Where("email = ?", userData.GetEmail()).First(&vUser).Error
// 		if err != nil {
// 			return err
// 		}
// 		if vUser == nil {
// 			return err
// 		}
// 		if userData.GetVerifyCode() != vUser.Code || time.Now().Unix() > vUser.ExpireTime.UnixNano() {
// 			return err
// 		}
// 		//
// 		//hash password
// 		userPassword, err := util.HashPassword(userData.GetPassword())
// 		if err != nil {
// 			return err
// 		}
// 		charset := "12345678"
// 		length := 8
// 		usr, _ := random.Random(length, charset, true)
// 		userName := "U" + usr
// 		currency := "USD"
// 		bankCard, err := util.GetBankCardNumber(userData.City)
// 		if err != nil {
// 			return err
// 		}
// 		user := &model.User{
// 			Username:       userName,
// 			Password:       userPassword,
// 			Email:          userData.Email,
// 			RegesiterIP:    "",
// 			LastLoginIP:    "",
// 			NationID:       userData.NationId,
// 			BankCardNumber: bankCard,
// 			City:           userData.City,
// 			Balance:        0,
// 			Currency:       currency,
// 			Type:           int8(userData.GenderType),
// 		}
// 		//
// 		fUser := &model.User{}
// 		data := server.Database.Sql.Model(&model.User{}).Where("email =?", userData.Email)
// 		err = data.First(&fUser).Error
// 		if err == gorm.ErrRecordNotFound {
// 			err := server.Database.Sql.Model(&model.User{}).Create(&user).Error
// 			if err != nil {
// 				return err
// 			}
// 			err = util.SaveUserBankCard(userData.City, user.ID, bankCard)
// 			if err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 		respUser = append(respUser, &pb.ResponseUser{
// 			User:      convert(user),
// 			CreatedAt: timestamppb.New(user.CreatedAt),
// 		})

// 		stream.Send(&pb.ResponseUserList{
// 			ResponseData: respUser,
// 		})

// 	}

// 	return nil
// }
