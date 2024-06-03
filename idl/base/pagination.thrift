namespace go sc_bff_base

//分页模型
struct Pagination {
    1: i64 pageNum, //当前页码
    2: i64 pageSize, //每页大小
    3: i64 totalItems, //总共条目数
}