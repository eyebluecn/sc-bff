package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/enums"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"net/http"
)

type EditorAuth struct {
}

func NewEditorAuth() *EditorAuth {
	return &EditorAuth{}
}

// 要求编辑已登录
func (receiver EditorAuth) Handle(c context.Context, appCtx *app.RequestContext) {
	readerVO := session.DefaultSession().FindLoginEditor(appCtx)
	if readerVO == nil {
		appCtx.JSON(enums.StatusCodeUnauthorized.HttpCode(), result.NewCodeWebResult(enums.StatusCodeUnauthorized, "编辑未登录，禁止访问"))
		appCtx.AbortWithStatus(http.StatusUnauthorized)
	}
}
