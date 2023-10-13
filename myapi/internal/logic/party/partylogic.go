package party

import (
	"context"

	"myapi/internal/svc"
	"myapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartyLogic {
	return &PartyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func NewPartyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartyLogic {
	return &PartyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartyLogic) Party(req *types.PartyReq) (resp *types.PartyResp, err error) {
	// todo: add your logic here and delete this line
	return &types.PartyResp{
		PartyId:   "12082",
		PartyName: "leo",
	}, nil
}
