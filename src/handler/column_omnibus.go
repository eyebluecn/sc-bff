package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/converter/vo2dto"
	"github.com/eyebluecn/sc-bff/src/infra/rpc"
	"github.com/eyebluecn/sc-bff/src/infra/session"
	"github.com/eyebluecn/sc-bff/src/model/result"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)
import "github.com/cloudwego/hertz/pkg/app"

type ColumnOmnibusHandler struct {
}

func NewColumnOmnibusHandler() *ColumnOmnibusHandler {
	return &ColumnOmnibusHandler{}
}

func (receiver ColumnOmnibusHandler) Handle(ctx context.Context, appCtx *app.RequestContext) (result.Response, error) {
	var err error
	var req sc_bff_api.ColumnOmnibusRequest
	err = appCtx.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}

	editorVO, err := session.DefaultSession().CheckLoginEditor(appCtx)
	if err != nil {
		return nil, err
	}
	operator := vo2dto.ConvertEditorOperator(editorVO)

	request := &sc_misc_api.ColumnOmnibusRequest{
		AuthorName:  req.AuthorName,
		ColumnName:  req.ColumnName,
		ColumnPrice: req.ColumnPrice,
		Operator:    operator,
	}

	richColumn, err := rpc.NewMiscCaller().ColumnOmnibus(ctx, request)
	if err != nil {
		return nil, err
	}

	klog.CtxInfof(ctx, "%v 创建专栏成功： %v", operator.Username, richColumn.Column.Id)

	return &sc_bff_api.ColumnOmnibusResponse{
		Code: 0,
		Msg:  "",
	}, nil

}
