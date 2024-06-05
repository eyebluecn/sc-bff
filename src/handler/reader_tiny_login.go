package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/constant"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/converter/vo2dto"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/model/result"
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/util"
)
import "github.com/cloudwego/hertz/pkg/app"

type ReaderTinyLoginHandler struct {
}

func NewReaderTinyLoginHandler() *ReaderTinyLoginHandler {
	return &ReaderTinyLoginHandler{}
}

func (receiver ReaderTinyLoginHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.ReaderTinyLoginRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	hlog.CtxInfof(ctx, "尝试tinyLogin")
	readerVo := session.DefaultSession().FindLoginReader(appCtx)
	if readerVo == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeLogin, "未登录，请先登录")
	} else {
		return &sc_bff_api.ReaderTinyLoginResponse{
			Data: vo2dto.ConvertReaderDTO(readerVo),
		}, nil
	}
}

func (receiver ReaderTinyLoginHandler) setCookieAndStorage(c context.Context, appCtx *app.RequestContext, readerVo *vo.ReaderVO) error {
	//登录成功后写入cookie.
	key := util.Uuid()
	//存储一年
	maxAge := 60 * 60 * 24 * 365
	appCtx.SetCookie(constant.READER_COOKIE_KEY, key, maxAge, "/", "", protocol.CookieSameSiteLaxMode, true, true)
	//cookie := appCtx .Response.Header.Get("Set-Cookie")
	//cookie == "user=hertz; max-age=1; domain=localhost; path=/; HttpOnly; secure; SameSite=Lax"

	//写到缓存中
	session.DefaultSession().StoreLoginReader(key, readerVo)

	return nil
}
