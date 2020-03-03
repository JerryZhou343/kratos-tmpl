package app

import (
	"github.com/google/wire"
	"github.com/mfslog/kratos-mvc/app/executor/cmd"
	"github.com/mfslog/kratos-mvc/app/executor/query"
	"github.com/mfslog/kratos-mvc/infra/repo"
)

var (
	AppProvider = wire.NewSet(cmd.NewAppCmd, query.NewAppQuery, repo.Provider)
)
