package api_conv

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为枚举
func ConvertPaymentStatus(status vo_model.PaymentStatus) sc_bff_api.PaymentStatus {
	return sc_bff_api.PaymentStatus(status)
}

// 转为DTO
func ConvertPaymentDTO(thing *vo_model.PaymentVO) *sc_bff_api.PaymentDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.PaymentDTO{
		ID:                 thing.ID,
		CreateTime:         util.Timestamp(thing.CreateTime),
		UpdateTime:         util.Timestamp(thing.UpdateTime),
		OrderNo:            thing.OrderNo,
		Method:             thing.Method,
		ThirdTransactionNo: thing.ThirdTransactionNo,
		Amount:             thing.Amount,
		Status:             ConvertPaymentStatus(thing.Status),
	}
}

// 转为DTO数组
func ConvertPaymentDTOS(things []*vo_model.PaymentVO) []*sc_bff_api.PaymentDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.PaymentDTO
	for _, item := range things {
		list = append(list, ConvertPaymentDTO(item))
	}
	return list
}
