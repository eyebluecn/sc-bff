package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/enums"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/infra/session"
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
		appCtx.JSON(enums.StatusCodeUnauthorized.HttpCode(), result.NewCodeWebResult(enums.StatusCodeUnauthorized, "读者未登录，禁止访问"))
		appCtx.AbortWithStatus(http.StatusUnauthorized)
	}
}
