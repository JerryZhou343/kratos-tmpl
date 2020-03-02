package cmd

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/mfslog/kratos-mvc/model/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppCmd struct {
	userRepo user.Repo
}

func NewAppCmd(userRepo user.Repo) *AppCmd {
	return &AppCmd{
		userRepo: userRepo,
	}
}

func (app *AppCmd) CreateUser(ctx context.Context, user *user.User) (ret *user.User, err error) {
	err = app.userRepo.CreateUser(ctx, user)
	if err != nil {
		log.Error("create user failed [%+v]", err)
		err = status.Error(codes.Internal, "create user failed")
	}
	ret = user
	return
}
