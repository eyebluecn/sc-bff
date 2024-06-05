package errs

type StatusCode int32

const (
	StatusCodeOk           StatusCode = 0    //成功
	StatusCodeLogin        StatusCode = 4001 //没有登录
	StatusCodeNotFound     StatusCode = 4002 //资源没找到
	StatusCodeBadRequest   StatusCode = 4003 //请求不合法
	StatusCodeUnauthorized StatusCode = 4004 //没有权限
	StatusCodeParamsError  StatusCode = 4005 //参数错误
	StatusCodeServer       StatusCode = 5000 //服务器内部出错
	StatusCodeUnknown      StatusCode = 5001 //服务器未知错误
)

// 状态码对应的http码
func (receiver StatusCode) HttpCode() int {
	switch receiver {
	case StatusCodeOk:
		return 200
	case StatusCodeLogin:
		return 400
	case StatusCodeNotFound:
		return 404
	case StatusCodeBadRequest:
		return 400
	case StatusCodeUnauthorized:
		return 401
	case StatusCodeParamsError:
		return 400
	case StatusCodeServer:
		return 500
	case StatusCodeUnknown:
		return 500
	default:
		return 500
	}
}
