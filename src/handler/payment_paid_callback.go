package handler

import (
	"context"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/converter/api_conv"
	"github.com/eyebluecn/sc-bff/src/infra/rpc"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)
import "github.com/cloudwego/hertz/pkg/app"

type PaymentPaidCallbackHandler struct {
}

func NewPaymentPaidCallbackHandler() *PaymentPaidCallbackHandler {
	return &PaymentPaidCallbackHandler{}
}

func (receiver PaymentPaidCallbackHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.PaymentPaidCallbackRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	request := &sc_misc_api.PaymentPaidCallbackRequest{
		OrderNo: req.OrderNo,
	}

	paymentVO, err := rpc.NewMiscCaller().PaymentPaidCallback(ctx, request)
	if err != nil {
		return nil, err
	}

	resp := &sc_bff_api.PaymentPaidCallbackResponse{
		Payment: api_conv.ConvertPaymentDTO(paymentVO),
		Code:    0,
		Msg:     "",
	}

	return resp, nil
}
