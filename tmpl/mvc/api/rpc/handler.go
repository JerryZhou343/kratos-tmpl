package rpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mfslog/kratos-mvc/app/executor/cmd"
	"github.com/mfslog/kratos-mvc/app/executor/query"
	userAssembler "github.com/mfslog/kratos-mvc/assembler/user"
	"github.com/mfslog/kratos-mvc/model/user"
	kratos "github.com/mfslog/kratos-mvc/proto"
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

func (h *Handler) GetUser(ctx context.Context, req *kratos.GetUserReq) (rsp *kratos.GetUserRsp, err error) {
	rsp = &kratos.GetUserRsp{}
	var (
		data *user.User
	)
	data, err = h.queryApp.GetUser(ctx, req.Id)
	rsp.User = userAssembler.UserAssembler.DO2DTO(data)
	return
}

func (h *Handler) CreateUser(ctx context.Context, req *kratos.CreateUserReq) (rsp *kratos.CreateUserRsp, err error) {
	rsp = &kratos.CreateUserRsp{}
	var (
		data *user.User
	)
	data, err = h.cmdApp.CreateUser(ctx, userAssembler.UserAssembler.DTO2DO(req.User))
	rsp.User = userAssembler.UserAssembler.DO2DTO(data)
	return
}
