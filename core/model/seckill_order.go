package model

import "time"

type SeckillOrder struct {
	Id        int64 `json:"id" gorm:"type:bigint(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserId    int64        `json:"user_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'用户ID';unique_index:u_uid_gid"`
	OrderId   int64        `json:"user_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'订单ID'"`
	Product   int64        `json:"user_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'产品ID';unique_index:u_uid_gid"`
}
