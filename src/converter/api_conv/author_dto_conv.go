package api_conv

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为DTO
func ConvertAuthorDTO(thing *vo_model.AuthorVO) *sc_bff_api.AuthorDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.AuthorDTO{
		ID:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Username:   thing.Username,
		Realname:   thing.Realname,
	}
}

// 转为DTO数组
func ConvertAuthorDTOS(things []*vo_model.AuthorVO) []*sc_bff_api.AuthorDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.AuthorDTO
	for _, item := range things {
		list = append(list, ConvertAuthorDTO(item))
	}
	return list
}
