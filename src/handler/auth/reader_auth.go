package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/model/result"
	"net/http"
)

type ReaderAuth struct {
}

func NewReaderAuth() *ReaderAuth {
	return &ReaderAuth{}
}

// 要求读者已登录
func (receiver ReaderAuth) Handle(c context.Context, appCtx *app.RequestContext) {
	readerVO := session.DefaultSession().FindLoginReader(appCtx)
	if readerVO == nil {
		//按照response返回
		appCtx.JSON(errs.StatusCodeUnauthorized.HttpCode(), result.NewCodeWebResult(errs.StatusCodeUnauthorized, "读者未登录，禁止访问"))
		appCtx.AbortWithStatus(http.StatusUnauthorized)
	}
}
