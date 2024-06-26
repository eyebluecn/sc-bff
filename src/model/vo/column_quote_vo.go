package vo

import (
	"github.com/eyebluecn/sc-bff/src/model/vo/enums"
	"time"
)

// 专栏 领域模型
type ColumnQuoteVO struct {
	ID         int64                   // 主键
	CreateTime time.Time               // 创建时间
	UpdateTime time.Time               // 更新时间
	ColumnID   int64                   // 专栏id
	EditorID   int64                   // 编辑id
	Price      int64                   // 价格（单位：分）
	Status     enums.ColumnQuoteStatus // 报价状态 0/1  未生效/已生效
}
