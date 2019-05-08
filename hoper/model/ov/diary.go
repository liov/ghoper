package ov

import (
	"time"
)

type Diary struct {
	ID          uint64    `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `gorm:"type:varchar(100)" json:"title"`
	DiaryBook   DiaryBook `json:"diaryBook"`
	DiaryBookID uint64    `json:"diary_book_id"`
	Content     string    `json:"content"`
	Mood        Mood      `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	//MoodName    string         `gorm:"type:varchar(20)" json:"mood_name"`
	Tags []Tag `gorm:"many2many:diary_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User User  `json:"user"`
	//UserID      uint64         `json:"user_id"`
	ImageUrl string `json:"image_url"` //封面
	ActionCount
	Permission uint8 `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence   uint8 `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status     uint8 `gorm:"type:smallint;default:0" json:"status"`     //状态
}

type DiaryBook struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	Title       string     `gorm:"type:varchar(100)" json:"title"`
	Diaries     []Diary    `json:"diaries"`
	Description string     `gorm:"type:varchar(500)" json:"description"` //描述
	Tags        []Tag      `gorm:"many2many:diary_book_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User        User       `json:"user"`
	//UserID      uint64             `json:"user_id"`
	ImageUrl string `gorm:"type:varchar(100)" json:"image_url"` //封面
	ActionCount
	Permission uint8 `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence   uint8 `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status     uint8 `gorm:"type:smallint;default:0" json:"status"`     //状态
}
