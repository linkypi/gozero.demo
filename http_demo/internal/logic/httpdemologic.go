package logic

import (
	"context"

	"http_demo/internal/svc"
	"http_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Http_demoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHttp_demoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Http_demoLogic {
	return &Http_demoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Http_demoLogic) Http_demo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
