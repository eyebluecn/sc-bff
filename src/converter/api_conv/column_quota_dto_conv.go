package api_conv

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为枚举
func ConvertColumnQuoteStatus(status vo_model.ColumnQuoteStatus) sc_bff_api.ColumnQuoteStatus {
	return sc_bff_api.ColumnQuoteStatus(status)
}

// 转为DTO
func ConvertColumnQuoteDTO(thing *vo_model.ColumnQuoteVO) *sc_bff_api.ColumnQuoteDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.ColumnQuoteDTO{
		ID:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		ColumnId:   thing.ColumnID,
		Editor:     thing.EditorID,
		Price:      thing.Price,
		Status:     ConvertColumnQuoteStatus(thing.Status),
	}
}

// 转为DTO数组
func ConvertColumnQuoteDTOS(things []*vo_model.ColumnQuoteVO) []*sc_bff_api.ColumnQuoteDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.ColumnQuoteDTO
	for _, item := range things {
		list = append(list, ConvertColumnQuoteDTO(item))
	}
	return list
}
