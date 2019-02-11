package model

import "time"

// Career 职业生涯
type Career struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Company   string     `json:"company"` //公司或组织
	Title     string     `json:"title"`   //职位
	UserID    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

// MaxCareerCompanyLen 公司或组织名称的最大长度
const MaxCareerCompanyLen = 200

// MaxCareerTitleLen 职位的最大长度
const MaxCareerTitleLen = 200
