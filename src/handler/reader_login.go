package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/constant"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/converter/api_conv"
	"github.com/eyebluecn/sc-bff/src/infra/rpc"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)
import "github.com/cloudwego/hertz/pkg/app"

type ReaderLoginHandler struct {
}

func NewReaderLoginHandler() *ReaderLoginHandler {
	return &ReaderLoginHandler{}
}

func (receiver ReaderLoginHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.ReaderLoginRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	request := &sc_misc_api.ReaderLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	readerVo, err := rpc.NewMiscCaller().ReaderLogin(ctx, request)
	if err != nil {
		return nil, err
	}

	hlog.CtxInfof(ctx, "%v 登录成功", readerVo.Username)

	//登录成功后，写入cookie和session
	_ = receiver.setCookieAndStorage(ctx, appCtx, readerVo)

	return &sc_bff_api.ReaderLoginResponse{
		Data: api_conv.ConvertReaderDTO(readerVo),
	}, nil
}

func (receiver ReaderLoginHandler) setCookieAndStorage(c context.Context, appCtx *app.RequestContext, readerVo *vo_model.ReaderVO) error {
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
