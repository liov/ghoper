package model

import (
	"time"
)

type ArticleComment struct {
	ID           uint64            `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time         `json:"created_at"`
	User         User              `json:"user"`
	UserID       uint64            `json:"user_id"`
	Content      string            `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64            `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User            `gorm:"many2many:article_comment_approve_user" json:"approve_users"`
	ArticleID    uint64            `json:"article_id"` //话题或投票的ID
	ParentID     uint64            `json:"parent_id"`  //直接父评论的ID
	RootID       uint64            `json:"root_id"`
	RecvUser     User              `gorm:"foreignkey:RecvUserID"json:"recv_user"`
	RecvUserID   uint64            `json:"recv_user_id"`
	SubComments  []*ArticleComment `gorm:"many2many:sub_article_comments;association_jointable_foreignkey:sub_id" json:"sub_comments"`
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type ArticleCommentApproveUser struct {
	CreatedAt        time.Time  `json:"created_at"`
	ArticleCommentID uint64     `gorm:"primary_key" json:"article_comment_id"`
	UserID           uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt        *time.Time `json:"-"`
	DeletedAt        *time.Time `sql:"index" json:"-"`
	Status           uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type MomentComment struct {
	ID           uint64           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	User         User             `json:"user"`
	UserID       uint64           `json:"user_id"`
	Content      string           `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64           `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User           `gorm:"many2many:moment_comment_approve_user" json:"approve_users"`
	MomentID     uint64           `json:"moment_id"` //瞬间ID
	ParentID     uint64           `json:"parent_id"` //直接父评论的ID
	RootID       uint64           `json:"root_id"`
	RecvUser     User             `gorm:"foreignkey:RecvUserID" json:"recv_user"`
	RecvUserID   uint64           `json:"recv_user_id"`
	SubComments  []*MomentComment `gorm:"many2many:sub_moment_comments;association_jointable_foreignkey:sub_id" json:"sub_comments"`
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type MomentCommentApproveUser struct {
	CreatedAt        time.Time  `json:"created_at"`
	ArticleCommentID uint64     `gorm:"primary_key" json:"moment_comment_id"`
	UserID           uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt        *time.Time `json:"-"`
	DeletedAt        *time.Time `sql:"index" json:"-"`
	Status           uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type DiaryComment struct {
	ID           uint64           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	User         User             `json:"user"`
	UserID       uint64           `json:"user_id"`
	Content      string           `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64           `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User           `gorm:"many2many:diary_comment_approve_user" json:"approve_users"`
	DiaryID      uint64           `json:"diary_id"`  //瞬间ID
	ParentID     uint64           `json:"parent_id"` //直接父评论的ID
	RootID       uint64           `json:"root_id"`
	RecvUser     User             `gorm:"foreignkey:RecvUserID" json:"recv_user"`
	RecvUserID   uint64           `json:"recv_user_id"`
	SubComments  []ArticleComment `gorm:"many2many:sub_diary_comment;association_jointable_foreignkey:sub_id" json:"sub_comments"`
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type DiaryCommentApproveUser struct {
	CreatedAt        time.Time  `json:"created_at"`
	ArticleCommentID uint64     `gorm:"primary_key" json:"diary_comment_id"`
	UserID           uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt        *time.Time `json:"-"`
	DeletedAt        *time.Time `sql:"index" json:"-"`
	Status           uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type DiaryBookComment struct {
	ID           uint64           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	User         User             `json:"user"`
	UserID       uint64           `json:"user_id"`
	Content      string           `gorm:"type:varchar(500)" json:"content"`
	ApproveCount uint64           `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User           `gorm:"many2many:diary_book_comment_approve_user" json:"approve_users"`
	DiaryBookID  uint64           `json:"diary_book_id"` //瞬间ID
	ParentID     uint64           `json:"parent_id"`     //直接父评论的ID
	RootID       uint64           `json:"root_id"`
	RecvUser     User             `gorm:"foreignkey:RecvUserID" json:"recv_user"`
	RecvUserID   uint64           `json:"recv_user_id"`
	SubComments  []ArticleComment `gorm:"many2many:sub_diary_book_comment;association_jointable_foreignkey:sub_id" json:"sub_comments"`
	//UpdatedAt *time.Time	`json:"updated_at"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type DiaryBookCommentApproveUser struct {
	CreatedAt        time.Time  `json:"created_at"`
	ArticleCommentID uint64     `gorm:"primary_key" json:"diary_book_comment_id"`
	UserID           uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt        *time.Time `json:"-"`
	DeletedAt        *time.Time `sql:"index" json:"-"`
	Status           uint8      `gorm:"type:smallint;default:0" json:"-"`
}

// MaxCommentLen 最大的评论长度
const MaxCommentLen = 5000
