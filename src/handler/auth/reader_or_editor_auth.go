package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/model/result"
	"net/http"
)

type ReaderOrEditorAuth struct {
}

func NewReaderOrEditorAuth() *ReaderOrEditorAuth {
	return &ReaderOrEditorAuth{}
}

// 要求读者或者编辑已登录
func (receiver ReaderOrEditorAuth) Handle(c context.Context, appCtx *app.RequestContext) {
	readerVO := session.DefaultSession().FindLoginReader(appCtx)
	editorVO := session.DefaultSession().FindLoginEditor(appCtx)
	if readerVO == nil && editorVO == nil {

		//按照response返回
		appCtx.JSON(errs.StatusCodeUnauthorized.HttpCode(), result.NewCodeWebResult(errs.StatusCodeUnauthorized, "读者或者编辑未登录，禁止访问"))
		appCtx.AbortWithStatus(http.StatusUnauthorized)
	}
}
