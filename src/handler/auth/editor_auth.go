package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/model/result"
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
		appCtx.JSON(errs.StatusCodeUnauthorized.HttpCode(), result.NewCodeWebResult(errs.StatusCodeUnauthorized, "编辑未登录，禁止访问"))
		appCtx.AbortWithStatus(http.StatusUnauthorized)
	}
}
