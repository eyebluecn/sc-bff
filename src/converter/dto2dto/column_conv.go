package dto2dto

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

// 转为枚举
func ConvertColumnStatus(status *int32) *sc_misc_api.ColumnStatus {
	if status == nil {
		return nil
	}
	res := sc_misc_api.ColumnStatus(*status)

	return &res
}
