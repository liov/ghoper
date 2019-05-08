package ov

import (
	"time"
)

type ArticleComment struct {
	ID           uint64           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	User         User             `json:"user"`
	UserID       uint64           `json:"user_id"`
	Content      string           `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64           `gorm:"default:0" json:"approve_count"` //点赞
	ArticleID    uint64           `json:"article_id"`                     //话题或投票的ID
	ParentID     uint64           `json:"parent_id"`                      //直接父评论的ID
	RootID       uint64           `json:"root_id"`
	RecvUser     User             `gorm:"foreignkey:RecvUserID" json:"recv_user"`
	RecvUserID   uint64           `json:"recv_user_id"`
	SubComments  []ArticleComment `gorm:"foreignkey:root_id" json:"sub_comments"`
	Sequence     uint8            `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
}

type MomentComment struct {
	ID           uint64          `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	User         User            `json:"user"`
	UserID       uint64          `json:"user_id"`
	Content      string          `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64          `gorm:"default:0" json:"approve_count"` //点赞
	MomentID     uint64          `json:"moment_id"`                      //瞬间ID
	ParentID     uint64          `json:"parent_id"`                      //直接父评论的ID
	RootID       uint64          `json:"root_id"`
	RecvUser     User            `gorm:"foreignkey:RecvUserID" json:"recv_user"`
	RecvUserID   uint64          `json:"recv_user_id"`
	SubComments  []MomentComment `gorm:"foreignkey:root_id" json:"sub_comments"`
	Sequence     uint8           `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
}

type DiaryComment struct {
	ID           uint64           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	User         User             `json:"user"`
	UserID       uint64           `json:"user_id"`
	Content      string           `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64           `gorm:"default:0" json:"approve_count"` //点赞
	DiaryID      uint64           `json:"diary_id"`                       //瞬间ID
	ParentID     uint64           `json:"parent_id"`                      //直接父评论的ID
	RootID       uint64           `json:"root_id"`
	RecvUser     User             `gorm:"foreignkey:RecvUserID" json:"recv_user"`
	RecvUserID   uint64           `json:"recv_user_id"`
	SubComments  []ArticleComment `gorm:"foreignkey:root_id" json:"sub_comments"`
	Sequence     uint8            `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
}

type DiaryBookComment struct {
	ID           uint64           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	User         User             `json:"user"`
	UserID       uint64           `json:"user_id"`
	Content      string           `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64           `gorm:"default:0" json:"approve_count"` //点赞
	DiaryBookID  uint64           `json:"diary_book_id"`                  //瞬间ID
	ParentID     uint64           `json:"parent_id"`                      //直接父评论的ID
	RootID       uint64           `json:"root_id"`
	RecvUser     User             `gorm:"foreignkey:RecvUserID" json:"recv_user"`
	RecvUserID   uint64           `json:"recv_user_id"`
	SubComments  []ArticleComment `gorm:"foreignkey:root_id" json:"sub_comments"`
	Sequence     uint8            `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
}
