package vo

import (
	"github.com/eyebluecn/sc-bff/src/model/vo/enums"
	"time"
)

// 专栏 领域模型
type ColumnVO struct {
	Id         int64              // 主键
	CreateTime time.Time          // 创建时间
	UpdateTime time.Time          // 更新时间
	Name       string             // 专栏名称
	AuthorId   int64              // 作者id
	Status     enums.ColumnStatus // 状态
}
