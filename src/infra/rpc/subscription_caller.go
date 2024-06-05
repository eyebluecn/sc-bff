package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/converter/dto2result"
	"github.com/eyebluecn/sc-bff/src/converter/dto2vo"
	"github.com/eyebluecn/sc-bff/src/converter/vo2dto"
	"github.com/eyebluecn/sc-bff/src/infra/rpc/config"
	"github.com/eyebluecn/sc-bff/src/model/result"
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

type SubscriptionCaller struct {
}

func NewSubscriptionCaller() SubscriptionCaller {
	return SubscriptionCaller{}
}

// 获取订阅分页列表。预期获得获取订阅分页列表，其余一律算错误。
// 如果err==nil，则ReaderVO!=nil
func (receiver SubscriptionCaller) SubscriptionPage(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest) ([]*vo.SubscriptionVO, *result.Pagination, error) {
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
	statusCode := dto2result.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != errs.StatusCodeOk {
		return nil, nil, errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}
	if response.Data == nil {
		return nil, nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: data is nil.")
	}
	if response.Pagination == nil {
		return nil, nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: pagination is nil.")
	}

	subscriptions := dto2vo.ConvertSubscriptions(response.Data)
	pagination := dto2result.ConvertSubscriptionPagination(response.Pagination)

	return subscriptions, pagination, nil
}

// 获取订阅分页列表。预期获得获取订阅分页列表，其余一律算错误。
// 如果err==nil，则ReaderVO!=nil
func (receiver SubscriptionCaller) SubscriptionRichPage(ctx context.Context, request *sc_subscription_api.SubscriptionRichPageRequest) ([]*vo.RichSubscriptionVO, *result.Pagination, error) {
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
	statusCode := dto2result.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != errs.StatusCodeOk {
		return nil, nil, errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}
	if response.Data == nil {
		return nil, nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: data is nil.")
	}
	if response.Pagination == nil {
		return nil, nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: pagination is nil.")
	}

	subscriptions := dto2vo.ConvertRichSubscriptionVOS(response.Data)
	pagination := dto2result.ConvertSubscriptionPagination(response.Pagination)

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
	statusCode := dto2result.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != errs.StatusCodeOk {
		return nil, errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}
	if response.Data == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: data is nil.")
	}

	orderVo := dto2vo.ConvertOrderVo(response.Data.OrderDTO)

	resp := &sc_bff_api.PrepareSubscribeDTO{
		OrderDTO:           vo2dto.ConvertOrderDTO(orderVo),
		ThirdTransactionNo: response.Data.ThirdTransactionNo,
		NonceStr:           response.Data.NonceStr,
	}

	return resp, nil
}
