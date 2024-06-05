package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	result2 "github.com/eyebluecn/sc-bff/src/model/result"
)
import "github.com/cloudwego/hertz/pkg/app"

type NotFoundHandler struct {
}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (receiver NotFoundHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result2.Response, error) {

	path := appCtx.GetRequest().URI().Path()
	pathStr := string(path)

	hlog.CtxErrorf(ctx, "没有http接口 %v", pathStr)
	return result2.NewCodeWebResult(errs.StatusCodeNotFound, fmt.Sprintf("找不到对应的方法 %v", pathStr)), nil
}
