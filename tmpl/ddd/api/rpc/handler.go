package rpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mfslog/kratos-ddd/app/executor/cmd"
	"github.com/mfslog/kratos-ddd/app/executor/query"
	userAssembler "github.com/mfslog/kratos-ddd/assembler/user"
	"github.com/mfslog/kratos-ddd/domain/aggregate/user"
	kratosddd "github.com/mfslog/kratos-ddd/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	cmdApp   *cmd.AppCmd
	queryApp *query.AppQuery
}

func NewHandler(cmdApp *cmd.AppCmd, queryApp *query.AppQuery) *Handler {
	return &Handler{
		cmdApp:   cmdApp,
		queryApp: queryApp,
	}
}

func (h *Handler) Ping(ctx context.Context, req *empty.Empty) (rsp *empty.Empty, err error) {
	rsp = &empty.Empty{}
	err = status.Error(codes.OK, "ok")
	return
}

func (h *Handler) GetUser(ctx context.Context, req *kratosddd.GetUserReq) (rsp *kratosddd.GetUserRsp, err error) {
	rsp = &kratosddd.GetUserRsp{}
	var (
		data *user.User
	)
	data, err = h.queryApp.GetUser(ctx, req.Id)
	rsp.User = userAssembler.UserAssembler.DO2DTO(data)
	return
}

func (h *Handler) CreateUser(ctx context.Context, req *kratosddd.CreateUserReq) (rsp *kratosddd.CreateUserRsp, err error) {
	rsp = &kratosddd.CreateUserRsp{}
	var (
		data *user.User
	)
	data, err = h.cmdApp.CreateUser(ctx, userAssembler.UserAssembler.DTO2DO(req.User))
	rsp.User = userAssembler.UserAssembler.DO2DTO(data)
	return
}
