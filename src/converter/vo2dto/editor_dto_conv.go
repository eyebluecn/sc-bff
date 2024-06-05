package vo2dto

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为DTO
func ConvertEditorDTO(thing *vo.EditorVO) *sc_bff_api.EditorDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.EditorDTO{
		ID:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}

// 转为DTO数组
func ConvertEditorDTOS(things []*vo.EditorVO) []*sc_bff_api.EditorDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.EditorDTO
	for _, item := range things {
		list = append(list, ConvertEditorDTO(item))
	}
	return list
}
