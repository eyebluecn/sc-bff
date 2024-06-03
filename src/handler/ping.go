package handler

import (
	"context"
	"fmt"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/util"
	"time"
)
import "github.com/cloudwego/hertz/pkg/app"

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (receiver PingHandler) Handle(c context.Context, ctx *app.RequestContext) (result.Response, error) {
	return result.NewWebResult(fmt.Sprintf("PONG! time = %v", util.FormatDateTime(time.Now()))), nil
}
