namespace go sc_bff_api


//作者模型
struct AuthorDTO {
    1: i64 id //id
    2: i64 createTime //创建时间
    3: i64 updateTime //编辑时间
    4: string username //昵称
    5: string realname //真名
}
