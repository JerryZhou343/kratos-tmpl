package grpc

import (
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/mfslog/kratos-ddd/infra/conf"
	pb "github.com/mfslog/kratos-ddd/proto"
)

// New new a grpc server.
func New(svc pb.KratosServer, config *conf.Config) (ws *warden.Server, err error) {
	ws = warden.NewServer(config.Grpc)
	pb.RegisterKratosServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
