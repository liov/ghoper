package controller

import "time"

type Category struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	Name           string     `json:"name"`
	Sequence       int        `json:"sequence"`  //同级别的分类可根据sequence的值来排序
	ParentID       int        `json:"parent_id"` //直接父分类的ID
	Articles       []Article  `json:"articles"`
	MomentCount    uint       `json:"moment_count"`
	ArticleCount   uint       `json:"article_count"`
	DiaryBookCount uint       `json:"diary_book_count"`
	DiaryCount     uint       `json:"diary_count"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`
	Status         uint8      `gorm:"type:smallint;default:0" json:"status"`
}
