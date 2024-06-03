package vo_conv

import (
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertSubscriptionStatus(status sc_subscription_api.SubscriptionStatus) vo_model.SubscriptionStatus {
	return vo_model.SubscriptionStatus(status)
}

// 转为VO
func ConvertSubscriptionVO(thing *sc_subscription_api.SubscriptionDTO) *vo_model.SubscriptionVO {
	if thing == nil {
		return nil
	}
	return &vo_model.SubscriptionVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		ReaderID:   thing.ReaderId,
		ColumnID:   thing.ColumnId,
		OrderID:    thing.OrderId,
		Status:     ConvertSubscriptionStatus(thing.Status),
	}
}

// 转为VO数组
func ConvertSubscriptions(subscriptionDTOS []*sc_subscription_api.SubscriptionDTO) []*vo_model.SubscriptionVO {
	if subscriptionDTOS == nil {
		return nil
	}
	var subscriptionVOS []*vo_model.SubscriptionVO
	for _, item := range subscriptionDTOS {
		subscriptionVOS = append(subscriptionVOS, ConvertSubscriptionVO(item))
	}
	return subscriptionVOS
}

// 转为枚举
func ConvertMiscSubscriptionStatus(status sc_misc_api.SubscriptionStatus) vo_model.SubscriptionStatus {
	return vo_model.SubscriptionStatus(status)
}

// 转为VO
func ConvertMiscSubscriptionVO(thing *sc_misc_api.SubscriptionDTO) *vo_model.SubscriptionVO {
	if thing == nil {
		return nil
	}
	return &vo_model.SubscriptionVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		ReaderID:   thing.ReaderId,
		ColumnID:   thing.ColumnId,
		OrderID:    thing.OrderId,
		Status:     ConvertMiscSubscriptionStatus(thing.Status),
	}
}

// 转为VO数组
func ConvertMiscSubscriptions(subscriptionDTOS []*sc_misc_api.SubscriptionDTO) []*vo_model.SubscriptionVO {
	if subscriptionDTOS == nil {
		return nil
	}
	var subscriptionVOS []*vo_model.SubscriptionVO
	for _, item := range subscriptionDTOS {
		subscriptionVOS = append(subscriptionVOS, ConvertMiscSubscriptionVO(item))
	}
	return subscriptionVOS
}

// 转为VO
func ConvertRichSubscriptionVO(thing *sc_subscription_api.RichSubscriptionDTO) *vo_model.RichSubscriptionVO {
	if thing == nil {
		return nil
	}
	return &vo_model.RichSubscriptionVO{
		Subscription: ConvertSubscriptionVO(thing.Subscription),
		Column:       ConvertColumnVO2(thing.Column),
		Reader:       ConvertReaderVO2(thing.Reader),
		Order:        ConvertOrderVo(thing.Order),
	}
}

// 转为VO数组
func ConvertRichSubscriptionVOS(things []*sc_subscription_api.RichSubscriptionDTO) []*vo_model.RichSubscriptionVO {
	if things == nil {
		return nil
	}
	var list []*vo_model.RichSubscriptionVO
	for _, item := range things {
		list = append(list, ConvertRichSubscriptionVO(item))
	}
	return list
}
