package auth

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"myapi/internal/logic"
	"myapi/internal/svc"
	"myapi/internal/types"
	"myapi/internal/utils"
)

type LoginLogic struct {
	logic.Logic
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	logic := &LoginLogic{}
	logic.Ctx = ctx
	logic.SvcCtx = svcCtx
	logic.Logger = logx.WithContext(ctx)
	return logic
}

//func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
//	return &LoginLogic{
//		Logger: logx.WithContext(ctx),
//		ctx:    ctx,
//		svcCtx: svcCtx,
//	}
//}

func (l LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	if req.Username == "admin" && req.Password == "password" {
		account := map[string]interface{}{
			"Username": req.Username,
			"Password": req.Password,
		}
		token, err := utils.GenJwtToken(account, l.SvcCtx.Config)
		if err != nil {
			return nil, err
		}
		return &types.LoginResp{
			Name:        req.Username,
			AccessToken: token.AccessToken,
		}, nil
	}

	return
}
