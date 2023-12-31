// Code generated by goctl. DO NOT EDIT.
// Source: rpc_demo.proto

package server

import (
	"context"

	"rpc_demo/internal/logic"
	"rpc_demo/internal/svc"
	"rpc_demo/rpc_demo"
)

type RpcDemoServer struct {
	svcCtx *svc.ServiceContext
	rpc_demo.UnimplementedRpcDemoServer
}

func NewRpcDemoServer(svcCtx *svc.ServiceContext) *RpcDemoServer {
	return &RpcDemoServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcDemoServer) Ping(ctx context.Context, in *rpc_demo.Request) (*rpc_demo.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
