include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/reader_model.thrift"
namespace go sc_bff_api

//读者获取自己信息 请求体
struct ReaderTinyLoginRequest {

}

//读者获取自己信息 响应体
struct ReaderTinyLoginResponse {
	1: reader_model.ReaderDTO data //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}


//读者登录 请求体
struct ReaderLoginRequest {
	1: string username //用户名
	2: string password //密码
}

//读者登录 响应体
struct ReaderLoginResponse {
	1: reader_model.ReaderDTO data //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}

//读者退出登录 请求体
struct ReaderLogoutRequest {
}

//读者退出登录 响应体
struct ReaderLogoutResponse {
    200: i32 code //状态码
    201: string msg //状态消息
}
