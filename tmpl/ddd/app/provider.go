package app

import (
	"github.com/google/wire"
	"github.com/mfslog/kratos-ddd/app/executor/cmd"
	"github.com/mfslog/kratos-ddd/app/executor/query"
	"github.com/mfslog/kratos-ddd/infra/repo"
)

var (
	AppProvider = wire.NewSet(cmd.NewAppCmd, query.NewAppQuery, repo.Provider)
)
