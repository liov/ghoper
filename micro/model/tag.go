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

type Mood struct {
	//ID        uint64 `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(20);primary_key" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	//DeletedAt *time.Time `sql:"index"`
	ExpressionURL string `gorm:"type:varchar(100)" json:"expression_url"`
	CreatedBy     User   `json:"created_by"`
	UserID        uint64 `json:"user_id"`
	/*	MomentCount    uint64   `json:"moment_count"`
		ArticleCount   uint64   `json:"article_count"`
		DiaryBookCount uint64   `json:"diary_book_count"`
		DiaryCount     uint64   `json:"diary_count"`*/
	Count     uint64     `gorm:"default:0" json:"count"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}
