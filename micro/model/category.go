package model

import "time"

type Category struct {
	ID             uint64     `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	Name           string     `json:"name"`
	ParentID       int        `json:"parent_id"` //直接父分类的ID
	Articles       []Article  `json:"articles"`
	Sequence       uint8      `gorm:"type:smallint;default:0" json:"sequence"` //同级别的分类可根据sequence的值来排序，置顶
	MomentCount    uint64     `json:"moment_count"`
	ArticleCount   uint64     `json:"article_count"`
	DiaryBookCount uint64     `json:"diary_book_count"`
	DiaryCount     uint64     `json:"diary_count"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`
	Status         uint8      `gorm:"type:smallint;default:0" json:"status"`
}
