package ov

import (
	"time"
)

type ArticleComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"` //点赞
	ContentType  int       `json:"content_type"`
	ArticleID    uint64    `json:"article_id"`                              //话题或投票的ID
	ParentID     uint64    `json:"parent_id"`                               //直接父评论的ID
	Sequence     uint8     `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	Status       uint8     `gorm:"type:smallint;default:0" json:"status"`
}

type MomentComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"`          //点赞
	MomentID     uint64    `json:"moment_id"`                               //瞬间ID
	ParentID     uint64    `json:"parent_id"`                               //直接父评论的ID
	Sequence     uint8     `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	Status       uint8     `gorm:"type:smallint;default:0" json:"status"`
}

type DiaryComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"`          //点赞
	DiaryID      uint64    `json:"diary_id"`                                //瞬间ID
	ParentID     uint64    `json:"parent_id"`                               //直接父评论的ID
	Sequence     uint8     `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	Status       uint8     `gorm:"type:smallint;default:0" json:"status"`
}

type DiaryBookComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"`          //点赞
	DiaryBookID  uint64    `json:"diary_book_id"`                           //瞬间ID
	ParentID     uint64    `json:"parent_id"`                               //直接父评论的ID
	Sequence     uint8     `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	Status       uint8     `gorm:"type:smallint;default:0" json:"status"`
}
