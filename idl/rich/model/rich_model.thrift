include "../../subscription/model/subscription_model.thrift"
include "../../column/model/column_model.thrift"
include "../../reader/model/reader_model.thrift"
include "../../order/model/order_model.thrift"
include "../../author/model/author_model.thrift"
include "../../column_quote/model/column_quote_model.thrift"
namespace go sc_bff_api


//富订阅模型
struct RichSubscriptionDTO {
    1: subscription_model.SubscriptionDTO subscription //订阅本身
    2: column_model.ColumnDTO column //专栏
    3: reader_model.ReaderDTO reader //读者
    4: order_model.OrderDTO order //订单
}



//富专栏模型
struct RichColumnDTO {
    1: column_model.ColumnDTO column //专栏本身
    2: author_model.AuthorDTO author //关联的作者
    3: column_quote_model.ColumnQuoteDTO columnQuote //关联的报价
    4: subscription_model.SubscriptionDTO subscription //关联的订阅情况
}