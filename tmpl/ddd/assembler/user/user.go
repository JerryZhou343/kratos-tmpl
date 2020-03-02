package user

import (
	"github.com/mfslog/kratos-ddd/domain/aggregate/user"
	kratosddd "github.com/mfslog/kratos-ddd/proto"
)

type userAssembler struct{}

var (
	UserAssembler = userAssembler{}
)

func (userAssembler) DO2DTO(src *user.User) *kratosddd.User {
	if src == nil {
		return nil
	}
	return &kratosddd.User{
		Id:   src.ID,
		Name: src.Name,
		Age:  src.Age,
	}
}

func (userAssembler) DTO2DO(src *kratosddd.User) *user.User {
	if src == nil {
		return nil
	}
	return &user.User{
		ID:   src.Id,
		Name: src.Name,
		Age:  src.Age,
	}
}
