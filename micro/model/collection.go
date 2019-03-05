package model

import (
	"time"
)

//收藏夹？像网易云一样可以收藏别人的歌单
type Collection struct {
	ID         uint        `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time   `json:"created_at"`
	User       *User       `json:"user"`
	UserID     uint        `json:"user_id"`
	Count      uint        `json:"count"`
	Articles   []Article   `gorm:"many2many:article_collection" json:"articles"`
	Moments    []Moment    `gorm:"many2many:moment_collection" json:"moments"`
	DiaryBooks []DiaryBook `gorm:"many2many:diary_book_collection" json:"diary_books"`
	Diarys     []Diary     `gorm:"many2many:diary_collection" json:"diarys"`
	UpdatedAt  *time.Time  `json:"updated_at"`
	DeletedAt  *time.Time  `sql:"index" json:"deleted_at"`
	Status     uint8       `gorm:"type:smallint;default:0" json:"status"`
}

type Like struct {
	ID         uint        `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time   `json:"created_at"`
	User       *User       `json:"user"`
	UserID     uint        `json:"user_id"`
	Count      uint        `json:"count"`
	Articles   []Article   `gorm:"many2many:article_like" json:"articles"`
	Moments    []Moment    `gorm:"many2many:moment_like" json:"moments"`
	DiaryBooks []DiaryBook `gorm:"many2many:diary_book_like" json:"diary_books"`
	Diaries    []Diary     `gorm:"many2many:diary_like" json:"diaries""`
	UpdatedAt  *time.Time  `json:"updated_at"`
	DeletedAt  *time.Time  `sql:"index" json:"deleted_at"`
	Status     uint8       `gorm:"type:smallint;default:0" json:"status"`
}
