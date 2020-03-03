package repo

import (
	"github.com/google/wire"
	userRepo "github.com/mfslog/kratos-mvc/infra/repo/demo"
)

var (
	Provider = wire.NewSet(userRepo.NewRepo)
)
