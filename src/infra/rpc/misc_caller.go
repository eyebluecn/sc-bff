package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/converter/dto2result"
	"github.com/eyebluecn/sc-bff/src/converter/dto2vo"
	"github.com/eyebluecn/sc-bff/src/infra/rpc/config"
	"github.com/eyebluecn/sc-bff/src/model/result"
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

type MiscCaller struct {
}

func NewMiscCaller() MiscCaller {
	return MiscCaller{}
}

// 读者登陆。预期获得登录对应的用户，其余一律算错误。
// 如果err==nil，则ReaderVO!=nil
func (receiver MiscCaller) ReaderLogin(ctx context.Context, request *sc_misc_api.ReaderLoginRequest) (*vo.ReaderVO, error) {
	response, err := config.MiscRpcClient.ReaderLogin(ctx, request)
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
	readerVo := dto2vo.ConvertReaderVo(response.Data)

	return readerVo, nil
}

// 小编登陆。预期获得登录对应的用户，其余一律算错误。
// 如果err==nil，则EditorVO!=nil
func (receiver MiscCaller) EditorLogin(ctx context.Context, request *sc_misc_api.EditorLoginRequest) (*vo.EditorVO, error) {
	response, err := config.MiscRpcClient.EditorLogin(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "Editor login failed: %v", err)
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
	editorVo := dto2vo.ConvertEditorVo(response.Data)

	return editorVo, nil
}

// 获取分页列表。预期获得获取分页列表，其余一律算错误。
// 如果err==nil，则ReaderVO!=nil
func (receiver MiscCaller) ColumnPage(ctx context.Context, request *sc_misc_api.RichColumnPageRequest) ([]*vo.RichColumnVO, *result.Pagination, error) {
	response, err := config.MiscRpcClient.RichColumnPage(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ColumnPage failed: %v", err)
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

	readerVos := dto2vo.ConvertRichColumnVOS(response.Data)
	pagination := receiver.convertPagination(response.Pagination)

	return readerVos, pagination, nil
}

// 演示接口 创建专栏等相关所有实体
func (receiver MiscCaller) ColumnOmnibus(ctx context.Context, request *sc_misc_api.ColumnOmnibusRequest) (*vo.RichColumnVO, error) {
	response, err := config.MiscRpcClient.ColumnOmnibus(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ColumnOmnibus failed: %v", err)
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
	if response.RichColumnDTO == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: data is nil.")
	}

	readerVos := dto2vo.ConvertRichColumnVO(response.RichColumnDTO)

	return readerVos, nil
}

// 支付成功回调方法，手动来触发。
func (receiver MiscCaller) PaymentPaidCallback(ctx context.Context, request *sc_misc_api.PaymentPaidCallbackRequest) (*vo.PaymentVO, error) {
	response, err := config.MiscRpcClient.PaymentPaidCallback(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ColumnOmnibus failed: %v", err)
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

	paymentVO := dto2vo.ConvertPaymentVO(response.Data)

	return paymentVO, nil
}
