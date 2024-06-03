package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/eyebluecn/sc-bff/src/common/enums"
	"github.com/eyebluecn/sc-bff/src/common/result"
)
import "github.com/cloudwego/hertz/pkg/app"

type NotFoundHandler struct {
}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (receiver NotFoundHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {

	path := appCtx.GetRequest().URI().Path()
	pathStr := string(path)

	hlog.CtxErrorf(ctx, "没有http接口 %v", pathStr)
	return result.NewCodeWebResult(enums.StatusCodeNotFound, fmt.Sprintf("找不到对应的方法 %v", pathStr)), nil
}
