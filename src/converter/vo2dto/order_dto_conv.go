package vo2dto

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/model/vo/enums"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为枚举
func ConvertOrderStatus(status enums.OrderStatus) sc_bff_api.OrderStatus {
	return sc_bff_api.OrderStatus(status)
}

// 转为DTO
func ConvertOrderDTO(thing *vo.OrderVO) *sc_bff_api.OrderDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.OrderDTO{
		ID:            thing.ID,
		CreateTime:    util.Timestamp(thing.CreateTime),
		UpdateTime:    util.Timestamp(thing.UpdateTime),
		No:            thing.No,
		ReaderID:      thing.ReaderID,
		ColumnID:      thing.ColumnID,
		ColumnQuoteID: thing.ColumnQuoteID,
		PaymentID:     thing.PaymentID,
		Price:         thing.Price,
		Status:        ConvertOrderStatus(thing.Status),
	}
}

// 转为DTO数组
func ConvertOrderDTOS(things []*vo.OrderVO) []*sc_bff_api.OrderDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.OrderDTO
	for _, item := range things {
		list = append(list, ConvertOrderDTO(item))
	}
	return list
}
