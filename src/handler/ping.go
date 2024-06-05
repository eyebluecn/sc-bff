package handler

import (
	"context"
	"fmt"
	result2 "github.com/eyebluecn/sc-bff/src/model/result"
	"github.com/eyebluecn/sc-bff/src/util"
	"time"
)
import "github.com/cloudwego/hertz/pkg/app"

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (receiver PingHandler) Handle(c context.Context, ctx *app.RequestContext) (result2.Response, error) {
	return result2.NewWebResult(fmt.Sprintf("PONG! time = %v", util.FormatDateTime(time.Now()))), nil
}
