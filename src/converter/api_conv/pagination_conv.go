package api_conv

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_base"
	"github.com/eyebluecn/sc-bff/src/model"
)

// 转为分页
func ConvertPagination(thing *model.Pagination) *sc_bff_base.Pagination {
	if thing == nil {
		return nil
	}
	return &sc_bff_base.Pagination{
		PageNum:    thing.PageNum,
		PageSize:   thing.PageSize,
		TotalItems: thing.TotalItems,
	}
}
