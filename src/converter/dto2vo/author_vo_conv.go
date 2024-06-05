package dto2vo

import (
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/util"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

func ConvertAuthorVo(thing *sc_misc_api.AuthorDTO) *vo.AuthorVO {
	if thing == nil {
		return nil
	}
	return &vo.AuthorVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Username:   thing.Username,
		Realname:   thing.Realname,
	}
}
