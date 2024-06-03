package result

import (
	"github.com/eyebluecn/sc-bff/src/common/enums"
	err2 "github.com/eyebluecn/sc-bff/src/common/errs"
)

type WebResult struct {
	Code enums.StatusCode `json:"code"`
	Msg  string           `json:"msg"`
	Data interface{}      `json:"data"`
}

func (p *WebResult) GetCode() (v int32) {
	return int32(p.Code)
}

func (p *WebResult) GetMsg() (v string) {
	return p.Msg
}

// 创建一个成功状态的WebResult
func NewWebResult(data interface{}) *WebResult {
	webResult := &WebResult{
		Code: enums.StatusCodeOk,
		Msg:  "成功",
		Data: data,
	}
	return webResult
}

// 创建一个失败状态的WebResult
func NewCodeWebResult(code enums.StatusCode, msg string) *WebResult {
	webResult := &WebResult{
		Code: code,
		Msg:  msg,
	}
	return webResult
}

// 创建一个失败状态的WebResult
func NewErrWebResult(err error) *WebResult {

	//尝试识别
	customErr := err2.Convert(err)
	if customErr == nil {
		return NewWebResult("OK")
	} else {
		return &WebResult{
			Code: customErr.Code,
			Msg:  customErr.Msg,
		}
	}
}
