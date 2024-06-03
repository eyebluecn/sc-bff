include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/rich_model.thrift"
namespace go sc_bff_api


//订阅列表 请求体
struct SubscriptionRichPageRequest {
	1: i64 pageNum //当前页码 1基
	2: i64 pageSize //每页大小

	3: optional i64 readerId //读者id
}

//订阅列表 响应体
struct SubscriptionRichPageResponse {
    1: SubscriptionRichPageData data //数据

    200: i32 code //状态码
    201: string msg //状态消息
}



//专栏详情 请求体
struct SubscriptionRichPageData {
	1: list<rich_model.RichSubscriptionDTO> items //数据信息
    2: pagination.Pagination pagination //分页指示器
}


//专栏列表 请求体
struct ColumnRichPageRequest {
	1: i64 pageNum //当前页码 1基
	2: i64 pageSize //每页大小
	4: optional string name //名称
	5: optional i64 authorId //作者id
	6: optional i32 status //状态
}

//专栏列表 响应体
struct ColumnRichPageResponse {
    1: ColumnRichPageData data //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}


//专栏详情 请求体
struct ColumnRichPageData {
	1: list<rich_model.RichColumnDTO> items //数据信息
    2: pagination.Pagination pagination //分页指示器
}
