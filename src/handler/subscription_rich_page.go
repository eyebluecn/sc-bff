package handler

import (
	"context"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/converter/api_conv"
	"github.com/eyebluecn/sc-bff/src/infra/rpc"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)
import "github.com/cloudwego/hertz/pkg/app"

type SubscriptionRichPageHandler struct {
}

func NewSubscriptionRichPageHandler() *SubscriptionRichPageHandler {
	return &SubscriptionRichPageHandler{}
}

func (receiver SubscriptionRichPageHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.SubscriptionRichPageRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	readerVO, err := session.DefaultSession().CheckLoginReader(appCtx)
	if err != nil {
		return nil, err
	}

	request := &sc_subscription_api.SubscriptionRichPageRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		ReaderId: util.ToInt64Ptr(readerVO.ID),
		Base:     nil,
	}

	richSubscriptionList, pagination, err := rpc.NewSubscriptionCaller().SubscriptionRichPage(ctx, request)
	if err != nil {
		return nil, err
	}

	return &sc_bff_api.SubscriptionRichPageResponse{
		Data: &sc_bff_api.SubscriptionRichPageData{
			Items:      api_conv.ConvertRichSubscriptionDTOS(richSubscriptionList),
			Pagination: api_conv.ConvertPagination(pagination),
		},
		Code: 0,
		Msg:  "",
	}, nil

}
