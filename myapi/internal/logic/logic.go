package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"myapi/internal/svc"
)

type Logic struct {
	logx.Logger
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		Logger: logx.WithContext(ctx),
		Ctx:    ctx,
		SvcCtx: svcCtx,
	}
}
