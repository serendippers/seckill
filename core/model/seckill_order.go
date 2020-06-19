package model

import "time"

type SeckillOrder struct {
	Id        int64 `json:"id" gorm:"type:bigint(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserId    int64      `json:"user_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'用户ID'"`
	OrderId   int64      `json:"order_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'订单ID'"`
	ProductId int64      `json:"product_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'产品ID'"`
}
