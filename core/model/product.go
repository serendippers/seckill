package model

import "time"

type Product struct {
	Id int `json:"id" gorm:"type:int(64)"`
	Name   string  `json:"name" gorm:"type:varchar(16);not null;default:'XXX';comment:'产品名称'"`
	Title  string  `json:"title" gorm:"type:varchar(64);default null;comment:'商品标题'"`
	Img    string  `json:"img" gorm:"type:varchar(64); DEFAULT NULL; COMMENT:'商品的图片'"`
	Detail string  `json:"detail" gorm:"type:longtext; DEFAULT NULL; COMMENT:'商品的详情介绍'"`
	Price  float32 `json:"price" gorm:"type:decimal(10,2); DEFAULT '0.00'; COMMENT:'商品单价'"`
	Stock  int     `json:"stock" gorm:"type:int(11); DEFAULT 0; COMMENT:'商品库存，-1表示没有限制'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
