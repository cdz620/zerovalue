package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerovalue/internal/logic"
	"zerovalue/internal/svc"
	"zerovalue/internal/types"
)

func ZerovalueHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewZerovalueLogic(r.Context(), svcCtx)
		resp, err := l.Zerovalue(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
