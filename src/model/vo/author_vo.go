package vo

import "time"

type AuthorVO struct {
	ID         int64     // 主键
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
	Username   string    // 用户名
	Realname   string    // 真实姓名
}
