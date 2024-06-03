package api_conv

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为枚举
func ConvertColumnStatus(status vo_model.ColumnStatus) sc_bff_api.ColumnStatus {
	return sc_bff_api.ColumnStatus(status)
}

// 转为DTO
func ConvertColumnDTO(thing *vo_model.ColumnVO) *sc_bff_api.ColumnDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.ColumnDTO{
		ID:         thing.Id,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Name:       thing.Name,
		AuthorId:   thing.AuthorId,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 转为DTO
func ConvertRichColumnDTO(thing *vo_model.RichColumnVO) *sc_bff_api.RichColumnDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.RichColumnDTO{
		Column:       ConvertColumnDTO(thing.Column),
		Author:       ConvertAuthorDTO(thing.Author),
		ColumnQuote:  ConvertColumnQuoteDTO(thing.ColumnQuote),
		Subscription: ConvertSubscriptionDTO(thing.Subscription),
	}
}

// 转为DTO数组
func ConvertRichColumnDTOS(things []*vo_model.RichColumnVO) []*sc_bff_api.RichColumnDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.RichColumnDTO
	for _, item := range things {
		list = append(list, ConvertRichColumnDTO(item))
	}
	return list
}
