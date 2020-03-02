package repo

import (
	"github.com/google/wire"
	userRepo "github.com/mfslog/kratos-mvc/infra/repo/user"
)

var (
	Provider = wire.NewSet(userRepo.NewRepo)
)
