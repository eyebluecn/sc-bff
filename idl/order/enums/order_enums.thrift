namespace go sc_bff_api

//状态
enum OrderStatus {
    CREATED = 0, //已创建
    PAID = 1, //已支付
    SUBSCRIBED = 2, //已订阅
    CLOSED = 3, //已关闭
    CANCELED = 4, //已取消
}
