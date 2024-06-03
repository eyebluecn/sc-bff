include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/payment_model.thrift"
namespace go sc_bff_api


// 请求体
struct PaymentPaidCallbackRequest {
    1: string orderNo //订单编号
}

// 响应体
struct PaymentPaidCallbackResponse {
    1: payment_model.PaymentDTO payment //数据信息

    200: i32 code //状态码
    201: string msg //状态消息
}
