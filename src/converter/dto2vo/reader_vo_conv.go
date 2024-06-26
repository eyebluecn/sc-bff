package dto2vo

import (
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

func ConvertReaderVo(thing *sc_misc_api.ReaderDTO) *vo.ReaderVO {
	if thing == nil {
		return nil
	}
	return &vo.ReaderVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}

func ConvertReaderVO2(thing *sc_subscription_api.ReaderDTO) *vo.ReaderVO {
	if thing == nil {
		return nil
	}
	return &vo.ReaderVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}
