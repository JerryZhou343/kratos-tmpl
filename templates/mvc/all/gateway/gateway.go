package gateway

import "github.com/mfslog/kratos-mvc/model/user"

type DDDClient interface {
	GetUser(id int64)(*user.User,error)
}