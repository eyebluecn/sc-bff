include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/column_model.thrift"
namespace go sc_bff_api

//请求体
struct ColumnOmnibusRequest {
	1: string authorName //作者名
	2: string columnName //专栏名
	3: i64 columnPrice //专栏价格
}

//响应体
struct ColumnOmnibusResponse {

    200: i32 code //状态码
    201: string msg //状态消息
}


//专栏详情 请求体
struct ColumnDetailRequest {
	1: i64 id //专栏Id
}

//专栏详情 响应体
struct ColumnDetailResponse {
	1: column_model.ColumnDTO data //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}
