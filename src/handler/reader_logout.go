package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/eyebluecn/sc-bff/src/common/constant"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	result2 "github.com/eyebluecn/sc-bff/src/model/result"
)
import "github.com/cloudwego/hertz/pkg/app"

type ReaderLogoutHandler struct {
}

func NewReaderLogoutHandler() *ReaderLogoutHandler {
	return &ReaderLogoutHandler{}
}

func (receiver ReaderLogoutHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result2.Response, error) {

	readerVO := session.DefaultSession().FindLoginReader(appCtx)
	if readerVO != nil {
		bytes := appCtx.Cookie(constant.READER_COOKIE_KEY)
		str := string(bytes)
		hlog.CtxInfof(ctx, "%v 退出登录 %v", readerVO.Username, str)
		session.DefaultSession().DelLoginReader(str)
	}

	return result2.NewWebResult("退出成功！"), nil

}
