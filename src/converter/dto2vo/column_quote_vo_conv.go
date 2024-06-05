package dto2vo

import (
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/model/vo/enums"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

// 转为枚举
func ConvertColumnQuoteStatus(thing sc_misc_api.ColumnQuoteStatus) enums.ColumnQuoteStatus {
	return enums.ColumnQuoteStatus(thing)
}

func ConvertColumnQuoteVo(thing *sc_misc_api.ColumnQuoteDTO) *vo.ColumnQuoteVO {
	if thing == nil {
		return nil
	}
	return &vo.ColumnQuoteVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		ColumnID:   thing.ColumnId,
		EditorID:   thing.EditorId,
		Price:      thing.Price,
		Status:     ConvertColumnQuoteStatus(thing.Status),
	}
}
