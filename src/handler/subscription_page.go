package handler

import (
	"context"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/converter/dto2dto"
	"github.com/eyebluecn/sc-bff/src/converter/vo2dto"
	"github.com/eyebluecn/sc-bff/src/infra/rpc"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/model/result"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)
import "github.com/cloudwego/hertz/pkg/app"

type SubscriptionPageHandler struct {
}

func NewSubscriptionPageHandler() *SubscriptionPageHandler {
	return &SubscriptionPageHandler{}
}

func (receiver SubscriptionPageHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.SubscriptionPageRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	readerVO, err := session.DefaultSession().CheckLoginReader(appCtx)
	if err != nil {
		return nil, err
	}

	request := &sc_subscription_api.SubscriptionPageRequest{
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
		ReaderId:  util.ToInt64Ptr(readerVO.ID),
		ColumnIds: req.ColumnIds,
		OrderId:   req.OrderId,
		Status:    dto2dto.ConvertSubscriptionStatus(req.Status),
	}

	subscriptionList, pagination, err := rpc.NewSubscriptionCaller().SubscriptionPage(ctx, request)
	if err != nil {
		return nil, err
	}

	return &sc_bff_api.SubscriptionPageResponse{
		Data:       vo2dto.ConvertSubscriptionDTOS(subscriptionList),
		Pagination: vo2dto.ConvertPagination(pagination),
		Code:       0,
		Msg:        "",
	}, nil

}
