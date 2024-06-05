package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/eyebluecn/sc-bff/src/common/constant"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	result2 "github.com/eyebluecn/sc-bff/src/model/result"
)
import "github.com/cloudwego/hertz/pkg/app"

type EditorLogoutHandler struct {
}

func NewEditorLogoutHandler() *EditorLogoutHandler {
	return &EditorLogoutHandler{}
}

func (receiver EditorLogoutHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result2.Response, error) {

	editorVO := session.DefaultSession().FindLoginEditor(appCtx)
	if editorVO != nil {
		bytes := appCtx.Cookie(constant.EDITOR_COOKIE_KEY)
		str := string(bytes)
		hlog.CtxInfof(ctx, "%v 退出登录 %v", editorVO.Username, str)
		session.DefaultSession().DelLoginEditor(str)
	}

	return result2.NewWebResult("退出成功！"), nil

}
