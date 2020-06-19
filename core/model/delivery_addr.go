package model

import "time"

type DeliveryAddr struct {
	Id           int64 `json:"id" gorm:"type:bigint(64)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	UserId       int64      `json:"user_id" gorm:"type:bigint(64);DEFAULT NULL; COMMENT:'用户ID';index:idx_user_id"`
	Consignee    string     `json:"consignee" gorm:"type:varchar(60);DEFAULT NULL; COMMENT:'收货人名字'"`
	ProvinceId   int        `json:"province_id" gorm:"type:int(4);NOT NULL;DEFAULT 0;COMMENT:'省份id'"`
	CityId       int        `json:"city_id" gorm:"type:int(4);NOT NULL;DEFAULT 0;COMMENT:'市id'"`
	DistrictId   int        `json:"district_id" gorm:"type:int(4);NOT NULL;DEFAULT 0;COMMENT:'区/县id'"`
	DeliveryAddr string     `json:"delivery_addr" gorm:"type:varchar(100);DEFAULT NULL;COMMENT:'收获地址'"`
	PhoneNumber  string     `json:"phone_number" gorm:"type:varchar(20);DEFAULT NULL; COMMENT:'电话'"`
}
