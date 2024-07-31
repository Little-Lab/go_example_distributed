package handler

import (
	"net/http"

	"bff/internal/logic"
	"bff/internal/svc"
	"bff/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func InviteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InviteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewInviteLogic(r.Context(), svcCtx)
		resp, err := l.Invite(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
