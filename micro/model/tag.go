package model

import "time"

type Tag struct {
	Name           string     `gorm:"type:varchar(10);primary_key" json:"name"`
	Description    string     `gorm:"type:varchar(100)" json:"description"`
	DeletedAt      *time.Time `sql:"index"`
	CreatedBy      User       `json:"created_by"`
	UserID         uint64     `json:"user_id"`
	MomentCount    uint64     `json:"moment_count"`
	ArticleCount   uint64     `json:"article_count"`
	DiaryBookCount uint64     `json:"diary_book_count"`
	DiaryCount     uint64     `json:"diary_count"`
	Count          uint64     `gorm:"default:0" json:"count"`
	Sequence       uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Status         uint8      `gorm:"type:smallint;default:0" json:"status"`
}
