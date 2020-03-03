package user

import (
	"github.com/mfslog/kratos-mvc/model/user"
	kratos "github.com/mfslog/kratos-mvc/proto"
)

type userAssembler struct{}

var (
	UserAssembler = userAssembler{}
)

func (userAssembler) DO2DTO(src *user.User) *kratos.User {
	if src == nil {
		return nil
	}
	return &kratos.User{
		Id:   src.ID,
		Name: src.Name,
		Age:  src.Age,
	}
}

func (userAssembler) DTO2DO(src *kratos.User) *user.User {
	if src == nil {
		return nil
	}
	return &user.User{
		ID:   src.Id,
		Name: src.Name,
		Age:  src.Age,
	}
}
