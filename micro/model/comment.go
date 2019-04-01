package model

import (
	"time"
)

type ArticleComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint64    `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	ContentType  int       `json:"content_type"`
	ArticleID    uint64    `json:"article_id"` //话题或投票的ID
	ParentID     uint64    `json:"parent_id"`  //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type MomentComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint64    `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	MomentID     uint64    `json:"moment_id"` //瞬间ID
	ParentID     uint64    `json:"parent_id"` //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type DiaryComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint64    `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	DiaryID      uint64    `json:"diary_id"`  //瞬间ID
	ParentID     uint64    `json:"parent_id"` //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type DiaryBookComment struct {
	ID           uint64    `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user"`
	UserID       uint64    `json:"user_id"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64    `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User    `gorm:"many2many:moment_approve_user" json:"approve_users"`
	DiaryBookID  uint64    `json:"diary_book_id"` //瞬间ID
	ParentID     uint64    `json:"parent_id"`     //直接父评论的ID
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
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
