package query

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/mfslog/kratos-mvc/model/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppQuery struct {
	userRepo user.Repo
}

func NewAppQuery(userRepo user.Repo) *AppQuery {
	return &AppQuery{
		userRepo: userRepo,
	}
}

func (app *AppQuery) GetUser(ctx context.Context, id int64) (ret *user.User, err error) {
	ret, err = app.userRepo.GetUser(ctx, id)
	if err != nil {
		log.Error("Get User failed [%+v]", err)
		err = status.Error(codes.Internal, "get demo failed")
	}
	return
}
