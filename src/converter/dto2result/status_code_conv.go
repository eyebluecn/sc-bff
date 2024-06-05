package dto2result

import (
	"github.com/eyebluecn/sc-bff/src/common/errs"
)

// 转为StatusCode
func ConvertStatusCode(statusCode int32) errs.StatusCode {
	return errs.StatusCode(statusCode)
}
