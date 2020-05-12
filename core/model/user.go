package model

import (
	"time"
)

type User struct {
	Id            int64 `json:"id" gorm:"type:bigint(64)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	Phone         string     `json:"phone" gorm:"type:varchar(13);NOT NULL;index:idx_phone;comment:'手机号码'"`
	Nickname      string     `json:"nickname" gorm:"type:varchar(255); NOT NULL; comment:'昵称'"`
	Password      string     `json:"password" gorm:"type: varchar(40); DEFAULT NULL; comment:'MD5(MD5(pass明文+固定salt) + salt)'"`
	Salt          string     `json:"salt" gorm:"type: varchar(10); DEFAULT NULL"`
	Head          string     `json:"head" gorm:"type:varchar(128); DEFAULT NULL; comment:'头像，云存储的ID'"`
	LoginCount    int        `json:"login_count" gorm:"type: int(11); DEFAULT 0; comment:'登录次数'"`
	LastLoginDate time.Time  `json:"last_login_date" gorm:"default:NULL;type:datetime; DEFAULT NULL; comment:'上次登陆时间'"`
}
