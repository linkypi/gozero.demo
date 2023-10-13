package handler

import (
	"myapi/internal/handler/auth"
	"time"

	"myapi/internal/handler/party"
	"myapi/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(routers *Routers, serverCtx *svc.ServiceContext) {

	group := routers.Group(
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
		rest.WithMaxBytes(1048576),
	)

	group.Post("/auth/login", auth.LoginHandler(serverCtx))

	// 开启 JWT 认证， 必须先经过JWT 认证后才会进入中间件
	group.AddOpt(rest.WithJwt(serverCtx.Config.Auth.AccessSecret))
	group.middleware = []rest.Middleware{serverCtx.AuthInterceptor}
	group.Post("/party/getByPage", party.PartyHandler(serverCtx))

}
