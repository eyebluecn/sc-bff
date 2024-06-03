package rpc

import (
	"github.com/eyebluecn/sc-bff/src/model"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
)

// 转为分页
func (receiver MiscCaller) convertPagination(pagination *sc_misc_base.Pagination) *model.Pagination {
	if pagination == nil {
		return nil
	}
	return &model.Pagination{
		PageNum:    pagination.PageNum,
		PageSize:   pagination.PageSize,
		TotalItems: pagination.TotalItems,
	}
}
