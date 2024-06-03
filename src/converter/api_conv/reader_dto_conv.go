package api_conv

import (
	"github.com/eyebluecn/sc-bff/idl_gen/sc_bff_api"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-bff/src/util"
)

// 转为DTO
func ConvertReaderDTO(thing *vo_model.ReaderVO) *sc_bff_api.ReaderDTO {
	if thing == nil {
		return nil
	}
	return &sc_bff_api.ReaderDTO{
		ID:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}

// 转为DTO数组
func ConvertReaderDTOS(things []*vo_model.ReaderVO) []*sc_bff_api.ReaderDTO {
	if things == nil {
		return nil
	}
	var list []*sc_bff_api.ReaderDTO
	for _, item := range things {
		list = append(list, ConvertReaderDTO(item))
	}
	return list
}
