include "../enums/payment_enums.thrift"
namespace go sc_bff_api


//支付单模型
struct PaymentDTO {
    1: i64 id //id
    2: i64 createTime //创建时间
    3: i64 updateTime //编辑时间

    4: string orderNo //订单编号
    5: string method //支付方式 ALIPAY/WEIXIN 支付宝/微信
    6: string thirdTransactionNo //支付平台订单号
    7: i64 amount //金额(单位：分)
    8: payment_enums.PaymentStatus status //支付状态 0/1/2  未支付/已支付/已关闭
}
