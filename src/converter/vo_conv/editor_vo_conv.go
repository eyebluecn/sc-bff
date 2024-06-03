package vo_conv

import (
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

func ConvertEditorVo(thing *sc_misc_api.EditorDTO) *vo_model.EditorVO {
	if thing == nil {
		return nil
	}
	return &vo_model.EditorVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}
