package api

import (
	"github.com/google/wire"
	"github.com/mfslog/kratos-ddd/api/rpc"
	"github.com/mfslog/kratos-ddd/app"
	kratos "github.com/mfslog/kratos-ddd/proto"
)

var (
	Provider = wire.NewSet(rpc.NewHandler, wire.Bind(new(kratos.KratosServer), new(*rpc.Handler)),
		app.AppProvider,
	)
)
