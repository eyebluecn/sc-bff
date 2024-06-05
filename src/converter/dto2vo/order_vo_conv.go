package dto2vo

import (
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/model/vo/enums"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertOrderStatus(status sc_subscription_api.OrderStatus) enums.OrderStatus {
	return enums.OrderStatus(status)
}

func ConvertOrderVo(thing *sc_subscription_api.OrderDTO) *vo.OrderVO {
	if thing == nil {
		return nil
	}
	return &vo.OrderVO{
		ID:            thing.Id,
		CreateTime:    util.ParseTimestamp(thing.CreateTime),
		UpdateTime:    util.ParseTimestamp(thing.UpdateTime),
		No:            thing.No,
		ReaderID:      thing.ReaderId,
		ColumnID:      thing.ColumnId,
		ColumnQuoteID: thing.ColumnQuoteId,
		PaymentID:     thing.PaymentId,
		Price:         thing.Price,
		Status:        ConvertOrderStatus(thing.Status),
	}
}
