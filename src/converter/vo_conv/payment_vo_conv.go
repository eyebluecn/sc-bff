package vo_conv

import (
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

// 转为枚举
func ConvertPaymentStatus(status sc_misc_api.PaymentStatus) vo_model.PaymentStatus {
	return vo_model.PaymentStatus(status)
}

// 转为VO
func ConvertPaymentVO(thing *sc_misc_api.PaymentDTO) *vo_model.PaymentVO {
	if thing == nil {
		return nil
	}
	return &vo_model.PaymentVO{
		ID:                 thing.Id,
		CreateTime:         util.ParseTimestamp(thing.CreateTime),
		UpdateTime:         util.ParseTimestamp(thing.UpdateTime),
		OrderNo:            thing.OrderNo,
		Method:             thing.Method,
		ThirdTransactionNo: thing.ThirdTransactionNo,
		Amount:             thing.Amount,
		Status:             ConvertPaymentStatus(thing.Status),
	}
}

// 转为VO数组
func ConvertPayments(things []*sc_misc_api.PaymentDTO) []*vo_model.PaymentVO {
	if things == nil {
		return nil
	}
	var subscriptionVOS []*vo_model.PaymentVO
	for _, item := range things {
		subscriptionVOS = append(subscriptionVOS, ConvertPaymentVO(item))
	}
	return subscriptionVOS
}