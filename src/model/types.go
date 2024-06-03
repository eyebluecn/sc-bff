package model

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/result"
)

// 标准的处理方法
type RouterFunc func(c context.Context, ctx *app.RequestContext) (result.Response, error)

// http接口处理器
type HttpHandler interface {
	Handle(c context.Context, appCtx *app.RequestContext) (result.Response, error)
}
