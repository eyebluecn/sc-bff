package result

// 标准的返回格式
type Response interface {
	GetCode() (v int32)
	GetMsg() (v string)
}

// 默认的返回体
type DefaultResponse struct {
	Code int32
	Msg  string
}

func (receiver DefaultResponse) GetCode() (v int32) {
	return receiver.Code
}
func (receiver DefaultResponse) GetMsg() (v string) {
	return receiver.Msg
}
