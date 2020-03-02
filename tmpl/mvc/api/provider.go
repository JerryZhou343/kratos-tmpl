package api

import (
	"github.com/google/wire"
	"github.com/mfslog/kratos-mvc/api/rpc"
	"github.com/mfslog/kratos-mvc/app"
	kratos "github.com/mfslog/kratos-mvc/proto"
)

var (
	Provider = wire.NewSet(rpc.NewHandler, wire.Bind(new(kratos.KratosServer), new(*rpc.Handler)),
		app.AppProvider,
	)
)
