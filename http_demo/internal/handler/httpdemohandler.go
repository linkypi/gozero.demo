package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"http_demo/internal/logic"
	"http_demo/internal/svc"
	"http_demo/internal/types"
)

func Http_demoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHttp_demoLogic(r.Context(), svcCtx)
		resp, err := l.Http_demo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
