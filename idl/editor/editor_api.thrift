include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/editor_model.thrift"
namespace go sc_bff_api

//小编获取自己信息 请求体
struct EditorTinyLoginRequest {

}

//小编获取自己信息 响应体
struct EditorTinyLoginResponse {
	1: editor_model.EditorDTO data //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}


//小编登录 请求体
struct EditorLoginRequest {
	1: string username //用户名
	2: string password //密码
}

//小编登录 响应体
struct EditorLoginResponse {
	1: editor_model.EditorDTO data //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}

//小编退出登录 请求体
struct EditorLogoutRequest {
}

//小编退出登录 响应体
struct EditorLogoutResponse {
    200: i32 code //状态码
    201: string msg //状态消息
}
