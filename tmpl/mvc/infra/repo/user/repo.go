package user

import (
	"context"
	"fmt"
	"github.com/mfslog/kratos-mvc/infra/conf"
	"github.com/mfslog/kratos-mvc/infra/repo/user/ent"
	entUser "github.com/mfslog/kratos-mvc/infra/repo/user/ent/user"
	"github.com/mfslog/kratos-mvc/model/user"
)

type repo struct {
	client *ent.Client
}

func NewRepo(conf *conf.Config) (user.Repo, error) {
	client, err := ent.Open("mysql", conf.SQL.DSN)
	fmt.Printf("dsn: %v\n", conf.SQL.DSN)
	if err != nil {
		return nil, err
	}
	return &repo{client: client}, nil
}

func (r *repo) GetUser(ctx context.Context, id int64) (*user.User, error) {
	ret, err := r.client.User.Query().Where(entUser.ID(int(id))).Only(ctx)

	return r.conv(ret), err
}

func (r *repo) CreateUser(ctx context.Context, data *user.User) error {
	ret, err := r.client.User.Create().SetAge(data.Age).SetName(data.Name).Save(ctx)
	data = r.conv(ret)
	return err
}

func (r *repo) conv(src *ent.User) *user.User {
	if src == nil {
		return nil
	}

	return &user.User{
		ID:   int64(src.ID),
		Name: src.Name,
		Age:  src.Age,
	}
}
