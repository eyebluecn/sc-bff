package dto2vo

import (
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/model/vo/enums"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertColumnStatus(status sc_misc_api.ColumnStatus) enums.ColumnStatus {
	return enums.ColumnStatus(status)
}

// 转为VO
func ConvertColumnVO(thing *sc_misc_api.ColumnDTO) *vo.ColumnVO {
	if thing == nil {
		return nil
	}
	return &vo.ColumnVO{
		Id:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Name:       thing.Name,
		AuthorId:   thing.AuthorId,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 转为枚举
func ConvertColumnStatus2(status sc_subscription_api.ColumnStatus) enums.ColumnStatus {
	return enums.ColumnStatus(status)
}

// 转为VO
func ConvertColumnVO2(thing *sc_subscription_api.ColumnDTO) *vo.ColumnVO {
	if thing == nil {
		return nil
	}
	return &vo.ColumnVO{
		Id:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Name:       thing.Name,
		AuthorId:   thing.AuthorId,
		Status:     ConvertColumnStatus2(thing.Status),
	}
}

// 转为VO数组
func ConvertColumns(things []*sc_misc_api.ColumnDTO) []*vo.ColumnVO {
	if things == nil {
		return nil
	}
	var subscriptionVOS []*vo.ColumnVO
	for _, item := range things {
		subscriptionVOS = append(subscriptionVOS, ConvertColumnVO(item))
	}
	return subscriptionVOS
}

// 转为VO
func ConvertRichColumnVO(thing *sc_misc_api.RichColumnDTO) *vo.RichColumnVO {
	if thing == nil {
		return nil
	}
	return &vo.RichColumnVO{
		Column:       ConvertColumnVO(thing.Column),
		Author:       ConvertAuthorVo(thing.Author),
		ColumnQuote:  ConvertColumnQuoteVo(thing.ColumnQuote),
		Subscription: ConvertMiscSubscriptionVO(thing.Subscription),
	}
}

// 转为VO数组
func ConvertRichColumnVOS(things []*sc_misc_api.RichColumnDTO) []*vo.RichColumnVO {
	if things == nil {
		return nil
	}
	var list []*vo.RichColumnVO
	for _, item := range things {
		list = append(list, ConvertRichColumnVO(item))
	}
	return list
}
