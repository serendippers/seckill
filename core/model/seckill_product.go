package model

import "time"

type SeckillProduct struct {
	Id        int64 `json:"id" gorm:"type:bigint(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	ProductId int64        `json:"product_id" gorm:"type:bigint(64); DEFAULT NULL; COMMENT:'商品Id'"`
	Price     float32    `json:"price" gorm:"type:decimal(10,2); DEFAULT '0.00'; COMMENT:'秒杀价'"`
	Stock     int        `json:"stock" gorm:"type:int(11); DEFAULT 0; COMMENT:'商品库存，-1表示没有限制'"`
	StartDate time.Time  `json:"last_login_date" gorm:"type:datetime; DEFAULT NULL; comment:'秒杀开始时间'"`
	EndDate   time.Time  `json:"last_login_date" gorm:"type:datetime; DEFAULT NULL; comment:'秒杀结束时间'"`
}
