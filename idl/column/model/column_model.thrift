include "../enums/column_enums.thrift"
namespace go sc_bff_api


//专栏模型
struct ColumnDTO {
    1: i64 id //id
    2: i64 createTime //创建时间
    3: i64 updateTime //编辑时间
    4: string name //名称
    5: i64 authorId //作者id
    6: column_enums.ColumnStatus status //状态
}

