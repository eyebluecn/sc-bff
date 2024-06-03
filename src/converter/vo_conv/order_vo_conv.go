package vo_conv

import (
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertOrderStatus(status sc_subscription_api.OrderStatus) vo_model.OrderStatus {
	return vo_model.OrderStatus(status)
}

func ConvertOrderVo(thing *sc_subscription_api.OrderDTO) *vo_model.OrderVO {
	if thing == nil {
		return nil
	}
	return &vo_model.OrderVO{
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
