package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/handler"
	"github.com/eyebluecn/sc-bff/src/handler/auth"
	result2 "github.com/eyebluecn/sc-bff/src/model/result"
	"github.com/eyebluecn/sc-bff/src/model/types"
)

type Router struct{}

func NewRouter() Router {
	return Router{}
}

// 注册所有的api方法
func (receiver Router) Register(h *server.Hertz) {

	//免登录接口
	publicGroup := h.Group("/api")
	publicGroup.GET("/ping", receiver.Wrap(handler.NewPingHandler()))
	publicGroup.POST("/reader/login", receiver.Wrap(handler.NewReaderLoginHandler()))
	publicGroup.POST("/reader/tiny/login", receiver.Wrap(handler.NewReaderTinyLoginHandler()))
	publicGroup.POST("/reader/logout", receiver.Wrap(handler.NewReaderLogoutHandler()))
	publicGroup.POST("/editor/login", receiver.Wrap(handler.NewEditorLoginHandler()))
	publicGroup.POST("/editor/tiny/login", receiver.Wrap(handler.NewEditorTinyLoginHandler()))
	publicGroup.POST("/editor/logout", receiver.Wrap(handler.NewEditorLogoutHandler()))

	//需要读者登录的接口
	readerGroup := h.Group("/api", auth.NewReaderAuth().Handle)
	readerGroup.GET("/subscription/page", receiver.Wrap(handler.NewSubscriptionPageHandler()))
	readerGroup.GET("/subscription/rich/page", receiver.Wrap(handler.NewSubscriptionRichPageHandler()))
	readerGroup.POST("/subscription/prepare", receiver.Wrap(handler.NewSubscriptionPrepareHandler()))
	readerGroup.POST("/payment/paid/callback", receiver.Wrap(handler.NewPaymentPaidCallbackHandler()))

	//需要编辑登录的接口
	editorGroup := h.Group("/api", auth.NewEditorAuth().Handle)
	editorGroup.POST("/column/omnibus", receiver.Wrap(handler.NewColumnOmnibusHandler()))

	//需要读者或者编辑登录的接口
	readerOrEditorGroup := h.Group("/api", auth.NewReaderOrEditorAuth().Handle)
	readerOrEditorGroup.GET("/column/rich/page", receiver.Wrap(handler.NewColumnRichPageHandler()))

	// 404配置
	h.NoRoute(receiver.Wrap(handler.NewNotFoundHandler()))
	h.NoMethod(receiver.Wrap(handler.NewNotFoundHandler()))

}

// 统一格式处理
func (receiver Router) Wrap(methodHandler types.HttpHandler) app.HandlerFunc {

	if methodHandler == nil {
		methodHandler = handler.NewNotFoundHandler()
	}

	var handlerFunc app.HandlerFunc = func(ctx context.Context, appCtx *app.RequestContext) {
		hlog.CtxInfof(ctx, "%v", appCtx.Request.URI())
		response, err := methodHandler.Handle(ctx, appCtx)
		customError := errs.Convert(err)
		if customError != nil {
			//用webResult作为response
			response = result2.NewErrWebResult(customError)
		}

		//可能返回体是nil
		if response == nil {
			response = result2.DefaultResponse{
				Code: 0,
				Msg:  "返回结果为空",
			}
		}

		statusCode := errs.StatusCode(response.GetCode())
		if statusCode != errs.StatusCodeOk {
			//出错了，打印错误
			hlog.CtxInfof(ctx, "error: %v uri: %v", response.GetMsg(), appCtx.Request.URI())
		}

		//按照response返回
		appCtx.JSON(statusCode.HttpCode(), response)

	}

	return handlerFunc
}
