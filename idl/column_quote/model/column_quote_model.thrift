namespace go sc_bff_api

include "../enums/column_quote_enums.thrift"


//专栏报价模型
struct ColumnQuoteDTO {
    1: i64 id //id
    2: i64 createTime //创建时间
    3: i64 updateTime //编辑时间
    4: i64 columnId //专栏id
    5: i64 editor //编辑id
    6: i64 price //价格
    7: column_quote_enums.ColumnQuoteStatus status //状态
}
