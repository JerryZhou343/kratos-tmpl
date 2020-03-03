package grpc

import (
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	pb "github.com/mfslog/kratos-mvc/proto"
)

// New new a grpc server.
func New(svc pb.KratosServer) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(&cfg)
	pb.RegisterKratosServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
