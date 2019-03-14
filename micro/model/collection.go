package model

import (
	"time"
)

type Favorites struct {
	ID          uint         `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	Name        string       `gorm:"type:varchar(20)" json:"name"`
	User        *User        `json:"user"`
	UserID      uint         `json:"user_id"`
	Count       uint         `json:"count"`
	Collections []Collection `gorm:"many2many:collection_favorites" json:"collections"`
	UpdatedAt   *time.Time   `json:"updated_at"`
	DeletedAt   *time.Time   `sql:"index" json:"deleted_at"`
	Status      uint8        `gorm:"type:smallint;default:0" json:"status"`
}

//收藏夹？像网易云一样可以收藏别人的歌单
type Collection struct {
	ID          uint        `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	RefID       uint        `json:"ref_id"`
	Kind        string      `gorm:"type:varchar(10)" json:"kind"`
	Favorites   []Favorites `json:"favorites"`
	FavoritesID uint        `json:"favorites_id"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `sql:"index" json:"deleted_at"`
	Status      uint8       `gorm:"type:smallint;default:0" json:"status"`
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

type ArticleCollection struct {
	ArticleID   uint         `json:"article_id"`
	Collections []Collection `json:"collections"`
}
