package vo_conv

import (
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

// 转为枚举
func ConvertColumnQuoteStatus(thing sc_misc_api.ColumnQuoteStatus) vo_model.ColumnQuoteStatus {
	return vo_model.ColumnQuoteStatus(thing)
}

func ConvertColumnQuoteVo(thing *sc_misc_api.ColumnQuoteDTO) *vo_model.ColumnQuoteVO {
	if thing == nil {
		return nil
	}
	return &vo_model.ColumnQuoteVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		ColumnID:   thing.ColumnId,
		EditorID:   thing.EditorId,
		Price:      thing.Price,
		Status:     ConvertColumnQuoteStatus(thing.Status),
	}
}
