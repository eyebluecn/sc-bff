package vo

// 专栏 领域模型
type RichColumnVO struct {
	Column       *ColumnVO       //专栏
	Author       *AuthorVO       //作者
	ColumnQuote  *ColumnQuoteVO  //作者
	Subscription *SubscriptionVO //订阅
}
