package model

import "time"

type OrderInfo struct {
	Id             int64 `json:"id" gorm:"type:bigint(64)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	UserId         int64      `json:"user_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'用户ID'"`
	ProductId      int64      `json:"product_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'产品ID'"`
	DeliveryAddrId int64      `json:"delivery_addr_id" gorm:"type:bigint(64);DEFAULT NULL;COMMENT:'收获地址ID'"`
	ProductName    string     `json:"product_name" gorm:"type:varchar(16); DEFAULT NULL; COMMENT:'冗余过来的商品名称'"`
	ProductCount   int        `json:"product_count" gorm:"type:int(11);DEFAULT 0; COMMENT:'商品数量'"`
	ProductPrice   float32    `json:"price" gorm:"type:decimal(10,2); DEFAULT '0.00'; COMMENT:'商品单价'"`
	OrderChannel   int        `json:"order_channel" gorm:"type:tinyint(4);DEFAULT 0; COMMENT:'1pc，2android，3ios'"`
	Status         int        `json:"status" gorm:"type:tinyint(4);DEFAULT 0;COMMENT:'订单状态，0新建未支付，1已支付，2已发货，3已收货，4已退款，5已完成'"`
	PayDate        time.Time  `json:"pay_date" gorm:"type:datetime;DEFAULT NULL; COMMENT: '支付时间'"`
}
