package vo_model

type SubscriptionStatus int32

const (
	SubscriptionStatusCreated  SubscriptionStatus = 0 //已创建
	SubscriptionStatusOk       SubscriptionStatus = 1 //已生效
	SubscriptionStatusDisabled SubscriptionStatus = 2 //已失效
)

type ColumnStatus int32

const (
	ColumnStatusNew      ColumnStatus = 0 //未发布
	ColumnStatusOK       ColumnStatus = 1 //已生效
	ColumnStatusDisabled ColumnStatus = 2 //已禁用
)

type ColumnQuoteStatus int32

const (
	ColumnQuoteStatusNew ColumnQuoteStatus = 0 //未生效
	ColumnQuoteStatusOK  ColumnQuoteStatus = 1 //已生效
)

// 订单状态
type OrderStatus int32

const (
	OrderStatusCreated    OrderStatus = 0 //已创建
	OrderStatusPaid       OrderStatus = 1 //已支付
	OrderStatusSubscribed OrderStatus = 2 //已订阅
	OrderStatusClosed     OrderStatus = 3 //已关闭
	OrderStatusCanceled   OrderStatus = 4 //已取消
)

type PaymentStatus int32

const (
	PaymentStatusUnpaid PaymentStatus = 0 //未支付
	PaymentStatusPaid   PaymentStatus = 1 //已支付
	PaymentStatusClosed PaymentStatus = 2 //已关闭
)
