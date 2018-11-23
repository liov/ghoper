package controller

import (
	"time"
)

type ArticleComment struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	UserID      uint      `json:"user_id"`
	Content     string    `gorm:"type:varchar(500)" json:"content"`
	HTMLContent string    `gorm:"type:varchar(500)" json:"html_content"`
	ContentType int       `json:"content_type"`
	ArticleID   uint      `json:"article_id"` //话题或投票的ID
	ParentID    uint      `json:"parent_id"`  //直接父评论的ID
}

type MomentComment struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint      `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	MomentID  uint      `json:"moment_id"` //瞬间ID
	ParentID  uint      `json:"parent_id"` //直接父评论的ID
}

type DiaryComment struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint      `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	DiaryID   uint      `json:"diary_id"`  //瞬间ID
	ParentID  uint      `json:"parent_id"` //直接父评论的ID
}

type DiaryBookComment struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	UserID      uint      `json:"user_id"`
	Content     string    `gorm:"type:varchar(500)" json:"content"`
	DiaryBookID uint      `json:"diary_book_id"` //瞬间ID
	ParentID    uint      `json:"parent_id"`     //直接父评论的ID
}
