package vo2dto

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/model/vo/enums"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为枚举
func ConvertSubscriptionStatus(status enums.SubscriptionStatus) sc_bff_api.SubscriptionStatus {
	return sc_bff_api.SubscriptionStatus(status)
}

// 转为DTO
func ConvertSubscriptionDTO(thing *vo.SubscriptionVO) *sc_bff_api.SubscriptionDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.SubscriptionDTO{
		ID:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		ReaderId:   thing.ReaderID,
		ColumnId:   thing.ColumnID,
		OrderId:    thing.OrderID,
		Status:     ConvertSubscriptionStatus(thing.Status),
	}
}

// 转为DTO数组
func ConvertSubscriptionDTOS(things []*vo.SubscriptionVO) []*sc_bff_api.SubscriptionDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.SubscriptionDTO
	for _, item := range things {
		list = append(list, ConvertSubscriptionDTO(item))
	}
	return list
}

// 转为DTO
func ConvertRichSubscriptionDTO(thing *vo.RichSubscriptionVO) *sc_bff_api.RichSubscriptionDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.RichSubscriptionDTO{
		Subscription: ConvertSubscriptionDTO(thing.Subscription),
		Column:       ConvertColumnDTO(thing.Column),
		Reader:       ConvertReaderDTO(thing.Reader),
		Order:        ConvertOrderDTO(thing.Order),
	}
}

// 转为DTO数组
func ConvertRichSubscriptionDTOS(things []*vo.RichSubscriptionVO) []*sc_bff_api.RichSubscriptionDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.RichSubscriptionDTO
	for _, item := range things {
		list = append(list, ConvertRichSubscriptionDTO(item))
	}
	return list
}
