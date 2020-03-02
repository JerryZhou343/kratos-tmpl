package di

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/mfslog/kratos-mvc/api/rpc"
	"github.com/mfslog/kratos-mvc/infra/conf"
	"time"
)

type App struct {
	svcHandler *rpc.Handler
	grpc       *warden.Server
	config     *conf.Config
}

func NewApp(svcHandler *rpc.Handler, g *warden.Server, config *conf.Config) (app *App, closeFunc func(), err error) {
	app = &App{
		svcHandler: svcHandler,
		grpc:       g,
		config:     config,
	}

	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := g.Shutdown(ctx); err != nil {
			log.Error("grpcSrv.Shutdown error(%v)", err)
		}
		cancel()
	}
	return
}
