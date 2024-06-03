namespace go sc_bff_api

//支付单状态
enum PaymentStatus {
    UNPAID  = 0, //未支付
    PAID = 1, //已支付
    CLOSED = 2, //已关闭
}
