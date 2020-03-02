// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/mfslog/kratos-ddd/api"
	"github.com/mfslog/kratos-ddd/infra/conf"
	"github.com/mfslog/kratos-ddd/infra/server/grpc"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(api.Provider, conf.NewConf, grpc.New, NewApp))
}
