package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/enums"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/infra/session"
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
		appCtx.JSON(enums.StatusCodeUnauthorized.HttpCode(), result.NewCodeWebResult(enums.StatusCodeUnauthorized, "读者或者编辑未登录，禁止访问"))
		appCtx.AbortWithStatus(http.StatusUnauthorized)
	}
}
