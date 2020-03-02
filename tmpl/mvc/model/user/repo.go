package user

import "context"

type Repo interface {
	GetUser(context.Context, int64) (*User, error)
	CreateUser(context.Context, *User) error
}
