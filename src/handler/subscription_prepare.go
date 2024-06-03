package handler

import (
	"context"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/infra/rpc"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)
import "github.com/cloudwego/hertz/pkg/app"

type SubscriptionPrepareHandler struct {
}

func NewSubscriptionPrepareHandler() *SubscriptionPrepareHandler {
	return &SubscriptionPrepareHandler{}
}

func (receiver SubscriptionPrepareHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.SubscriptionPrepareRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	readerVO, err := session.DefaultSession().CheckLoginReader(appCtx)
	if err != nil {
		return nil, err
	}

	request := &sc_subscription_api.SubscriptionPrepareRequest{
		ColumnId:  req.ColumnId,
		PayMethod: req.PayMethod,
		ReaderId:  readerVO.ID,
	}

	prepareSubscribeDTO, err := rpc.NewSubscriptionCaller().SubscriptionPrepare(ctx, request)
	if err != nil {
		return nil, err
	}

	resp := &sc_bff_api.SubscriptionPrepareResponse{
		Data: prepareSubscribeDTO,
		Code: 0,
		Msg:  "",
	}

	return resp, nil
}
