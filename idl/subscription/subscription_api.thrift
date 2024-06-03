include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/subscription_model.thrift"
include "../order/model/order_model.thrift"
namespace go sc_bff_api

//订阅下单 请求体
struct SubscriptionPrepareRequest {
	1: i64 columnId //专栏id
	2: string payMethod //支付方式
}

//订阅下单 响应体
struct SubscriptionPrepareResponse {
	1: PrepareSubscribeDTO data //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}

//下单结构体信息
struct PrepareSubscribeDTO {
    1: order_model.OrderDTO orderDTO //相关的订单
    2: string thirdTransactionNo //支付的一些token及信息
    3: string nonceStr //支付时候的验证信息等
}

//订阅列表 请求体
struct SubscriptionPageRequest {
	1: i64 pageNum //当前页码 1基
	2: i64 pageSize //每页大小
	4: optional list<i64> columnIds //专栏id
	5: optional i64 orderId //订单id
	6: optional i32 status //状态
}

//订阅列表 响应体
struct SubscriptionPageResponse {
	1: list<subscription_model.SubscriptionDTO> data //数据信息
    2: pagination.Pagination pagination //分页指示器

    200: i32 code //状态码
    201: string msg //状态消息
}
