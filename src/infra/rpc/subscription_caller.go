package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/config"
	"github.com/eyebluecn/sc-bff/src/common/enums"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/converter/api_conv"
	"github.com/eyebluecn/sc-bff/src/converter/model_conv"
	"github.com/eyebluecn/sc-bff/src/converter/vo_conv"
	"github.com/eyebluecn/sc-bff/src/model"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

type SubscriptionCaller struct {
}

func NewSubscriptionCaller() SubscriptionCaller {
	return SubscriptionCaller{}
}

// 获取订阅分页列表。预期获得获取订阅分页列表，其余一律算错误。
// 如果err==nil，则ReaderVO!=nil
func (receiver SubscriptionCaller) SubscriptionPage(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest) ([]*vo_model.SubscriptionVO, *model.Pagination, error) {
	response, err := config.SubscriptionRpcClient.SubscriptionPage(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "SubscriptionPage failed: %v", err)
		return nil, nil, err
	}
	if response == nil {
		klog.CtxErrorf(ctx, "Failed: response is nil")
		return nil, nil, errs.Errorf("response is nil")
	}
	if response.BaseResp == nil {
		klog.CtxErrorf(ctx, "Failed: response.BaseResp is nil")
		return nil, nil, errs.Errorf("response is nil")
	}
	statusCode := model_conv.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != enums.StatusCodeOk {
		return nil, nil, errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}
	if response.Data == nil {
		return nil, nil, errs.CodeErrorf(enums.StatusCodeNotFound, "Failed: data is nil.")
	}
	if response.Pagination == nil {
		return nil, nil, errs.CodeErrorf(enums.StatusCodeNotFound, "Failed: pagination is nil.")
	}

	subscriptions := vo_conv.ConvertSubscriptions(response.Data)
	pagination := model_conv.ConvertSubscriptionPagination(response.Pagination)

	return subscriptions, pagination, nil
}

// 获取订阅分页列表。预期获得获取订阅分页列表，其余一律算错误。
// 如果err==nil，则ReaderVO!=nil
func (receiver SubscriptionCaller) SubscriptionRichPage(ctx context.Context, request *sc_subscription_api.SubscriptionRichPageRequest) ([]*vo_model.RichSubscriptionVO, *model.Pagination, error) {
	response, err := config.SubscriptionRpcClient.SubscriptionRichPage(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "SubscriptionRichPage failed: %v", err)
		return nil, nil, err
	}
	if response == nil {
		klog.CtxErrorf(ctx, "Failed: response is nil")
		return nil, nil, errs.Errorf("response is nil")
	}
	if response.BaseResp == nil {
		klog.CtxErrorf(ctx, "Failed: response.BaseResp is nil")
		return nil, nil, errs.Errorf("response is nil")
	}
	statusCode := model_conv.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != enums.StatusCodeOk {
		return nil, nil, errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}
	if response.Data == nil {
		return nil, nil, errs.CodeErrorf(enums.StatusCodeNotFound, "Failed: data is nil.")
	}
	if response.Pagination == nil {
		return nil, nil, errs.CodeErrorf(enums.StatusCodeNotFound, "Failed: pagination is nil.")
	}

	subscriptions := vo_conv.ConvertRichSubscriptionVOS(response.Data)
	pagination := model_conv.ConvertSubscriptionPagination(response.Pagination)

	return subscriptions, pagination, nil
}

func (receiver SubscriptionCaller) SubscriptionPrepare(ctx context.Context,
	request *sc_subscription_api.SubscriptionPrepareRequest) (*sc_bff_api.PrepareSubscribeDTO, error) {
	response, err := config.SubscriptionRpcClient.SubscriptionPrepare(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "Reader login failed: %v", err)
		return nil, err
	}
	if response == nil {
		klog.CtxErrorf(ctx, "Failed: response is nil")
		return nil, errs.Errorf("response is nil")
	}
	if response.BaseResp == nil {
		klog.CtxErrorf(ctx, "Failed: response.BaseResp is nil")
		return nil, errs.Errorf("response is nil")
	}
	statusCode := model_conv.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != enums.StatusCodeOk {
		return nil, errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}
	if response.Data == nil {
		return nil, errs.CodeErrorf(enums.StatusCodeNotFound, "Failed: data is nil.")
	}

	orderVo := vo_conv.ConvertOrderVo(response.Data.OrderDTO)

	resp := &sc_bff_api.PrepareSubscribeDTO{
		OrderDTO:           api_conv.ConvertOrderDTO(orderVo),
		ThirdTransactionNo: response.Data.ThirdTransactionNo,
		NonceStr:           response.Data.NonceStr,
	}

	return resp, nil
}
