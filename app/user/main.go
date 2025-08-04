package main

import (
	"context"

	"gorm.io/gorm"
	gormplugin "trpc.group/trpc-go/trpc-database/gorm"

	"github.com/kickastone/trpc-demo/app/user/model"
	user "github.com/kickastone/trpc-demo/proto/user"
	"trpc.group/trpc-go/trpc-go"
)

type UserServiceImpl struct{}

var gormDB *gorm.DB

func (s *UserServiceImpl) GetAccountByUserName(ctx context.Context, req *user.GetAccountByUserNameRequest) (*user.GetAccountByUserNameResponse, error) {

	var userInfo model.UserInfo
	err := gormDB.Where("nickname = ?", req.Username).First(&userInfo).Error
	if err != nil {
		return &user.GetAccountByUserNameResponse{
			ErrCode: -1,
			ErrMsg:  "failed to get user info",
		}, nil
	}

	return &user.GetAccountByUserNameResponse{
		ErrCode: 0,
		ErrMsg:  "success",
		Data: &user.GetAccountByUserNameResponse_Data{
			UserInfo: &user.UserInfo{
				UserId:       userInfo.UUID,
				Username:     userInfo.Nickname,
				PasswordHash: userInfo.Password,
				Nickname:     userInfo.Nickname,
				CreateTsSec:  userInfo.CreatedAt.Unix(),
			},
		},
	}, nil
}

func main() {
	s := trpc.NewServer()
	gormDB, _ = gormplugin.NewClientProxy("trpc.mysql.user.User")
	user.RegisterUserService(s, &UserServiceImpl{})
	s.Serve()
}
