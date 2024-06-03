package handler

import (
	"context"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/common/result"
	"github.com/eyebluecn/sc-bff/src/converter/api2vo_conv"
	"github.com/eyebluecn/sc-bff/src/converter/api_conv"
	"github.com/eyebluecn/sc-bff/src/converter/vo2vo_conv"
	"github.com/eyebluecn/sc-bff/src/infra/rpc"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
)
import "github.com/cloudwego/hertz/pkg/app"

type ColumnRichPageHandler struct {
}

func NewColumnRichPageHandler() *ColumnRichPageHandler {
	return &ColumnRichPageHandler{}
}

func (receiver ColumnRichPageHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.ColumnRichPageRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	//解析出当前的操作者
	var operator *sc_misc_base.Operator
	readerVO := session.DefaultSession().FindLoginReader(appCtx)
	if readerVO != nil {
		operator = vo2vo_conv.ConvertReaderOperator(readerVO)
	} else {
		editorVO := session.DefaultSession().FindLoginEditor(appCtx)
		if editorVO != nil {
			operator = vo2vo_conv.ConvertEditorOperator(editorVO)
		}
	}

	request := &sc_misc_api.RichColumnPageRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Name:     nil,
		AuthorId: nil,
		Status:   api2vo_conv.ConvertColumnStatus(req.Status),
		Operator: operator,
	}

	columnList, pagination, err := rpc.NewMiscCaller().ColumnPage(ctx, request)
	if err != nil {
		return nil, err
	}

	return &sc_bff_api.ColumnRichPageResponse{
		Data: &sc_bff_api.ColumnRichPageData{
			Items:      api_conv.ConvertRichColumnDTOS(columnList),
			Pagination: api_conv.ConvertPagination(pagination),
		},
		Code: 0,
		Msg:  "",
	}, nil

}
