// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/mfslog/kratos-mvc/api"
	"github.com/mfslog/kratos-mvc/infra/conf"
	"github.com/mfslog/kratos-mvc/infra/server/grpc"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(api.Provider, conf.NewConf, grpc.New, NewApp))
}
