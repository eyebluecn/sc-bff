package vo_model

// 订阅 领域模型
type RichSubscriptionVO struct {
	Subscription *SubscriptionVO
	Column       *ColumnVO
	Reader       *ReaderVO
	Order        *OrderVO
}
