package auth

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"myapi/internal/svc"
	"myapi/internal/types"
	"myapi/internal/utils"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	if req.Username == "admin" && req.Password == "password" {
		account := map[string]interface{}{
			"Username": req.Username,
			"Password": req.Password,
		}
		token, err := utils.GenJwtToken(account, l.svcCtx.Config)
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
