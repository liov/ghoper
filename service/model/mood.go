package model

import "time"

type Mood struct {
	//ID        uint `gorm:"primary_key"`
	Name           string `gorm:"type:varchar(20);primary_key" json:"name"`
	Description string  `gorm:"type:varchar(100)" json:"description"`
	//DeletedAt *time.Time `sql:"index"`
	ExpressionURL  string `gorm:"type:varchar(100)" json:"expression_url"`
	CreatedBy      User   `json:"created_by"`
	UserID         uint   `json:"user_id"`
/*	MomentCount    uint   `json:"moment_count"`
	ArticleCount   uint   `json:"article_count"`
	DiaryBookCount uint   `json:"diary_book_count"`
	DiaryCount     uint   `json:"diary_count"`*/
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt *time.Time	`json:"updated_at"`
	Count	uint	`gorm:"default:0" json:"count"`
	Status         uint8  `gorm:"type:smallint;default:0" json:"status"`
}
