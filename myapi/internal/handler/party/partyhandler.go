package party

import (
	"github.com/zeromicro/go-zero/core/logx"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myapi/internal/logic/party"
	"myapi/internal/svc"
	"myapi/internal/types"

	_ "github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

// PartyHandler
func PartyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartyReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, errors.New(505, "参数异常"))
			logx.Error(err)
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(505, "参数异常"))
			//httpx.ErrorCtx(r.Context(), w, err)
			//fmt.Println(err)
			return
		}

		l := party.NewPartyLogic(r.Context(), svcCtx)
		resp, err := l.Party(&req)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			//httpx.OkJsonCtx(r.Context(), w, resp)
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
