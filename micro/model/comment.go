package model

import (
	"time"
)

type ArticleComment struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint      `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint      `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	ContentType  int       `json:"content_type"`
	ArticleID    uint      `json:"article_id"` //话题或投票的ID
	ParentID     uint      `json:"parent_id"`  //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

type MomentComment struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint      `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint      `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	MomentID     uint      `json:"moment_id"` //瞬间ID
	ParentID     uint      `json:"parent_id"` //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

type DiaryComment struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint      `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint      `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	DiaryID      uint      `json:"diary_id"`  //瞬间ID
	ParentID     uint      `json:"parent_id"` //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

type DiaryBookComment struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint      `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint      `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	DiaryBookID  uint      `json:"diary_book_id"` //瞬间ID
	ParentID     uint      `json:"parent_id"`     //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

const (
	// CommentSourceArticle 对话题进行评论
	CommentSourceArticle = "article"

	// CommentSourceVote 对投票进行评论
	CommentSourceVote = "vote"
)

const (
	// CommentVerifying 审核中
	CommentVerifying = 1

	// CommentVerifySuccess 审核通过
	CommentVerifySuccess = 2

	// CommentVerifyFail 审核未通过
	CommentVerifyFail = 3
)

// MaxCommentLen 最大的评论长度
const MaxCommentLen = 5000
