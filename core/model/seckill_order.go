package model

import "time"

type SeckillOrder struct {
	Id        int `json:"id" gorm:"type:int(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserId    int        `json:"user_id" gorm:"type:int(64);DEFAULT NULL; COMMENT:'用户ID';unique_index:u_uid_gid"`
	OrderId   int        `json:"user_id" gorm:"type:int(64);DEFAULT NULL; COMMENT:'订单ID'"`
	Product   int        `json:"user_id" gorm:"type:int(64);DEFAULT NULL; COMMENT:'产品ID';unique_index:u_uid_gid"`
}