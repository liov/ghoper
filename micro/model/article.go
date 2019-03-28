package model

import (
	"time"
)

type Article struct {
	ID          uint64           `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time        `json:"created_at"`
	Title       string           `gorm:"type:varchar(100)" json:"title"`
	Intro       string           `gorm:"type:varchar(100)" json:"intro"`
	Abstract    string           `gorm:"type:varchar(200)" json:"abstract"`
	Content     string           `gorm:"type:text" json:"content"`
	HTMLContent string           `gorm:"type:text" json:"html_content"`
	ContentType int              `json:"content_type"`                                 //文本类型
	ImageUrl    string           `gorm:"type:varchar(100)" json:"image_url"`           //封面
	Categories  []Category       `gorm:"many2many:article_category" json:"categories"` //分类
	Tags        []Tag            `gorm:"many2many:article_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User        User             `json:"user"`
	UserID      uint64           `json:"user_id"`
	Comments    []ArticleComment `gorm:"ForeignKey:ArticleID" json:"comments"` //评论
	ActionCount
	ApproveUsers  []User     `gorm:"many2many:article_approve_user" json:"approve_users"`
	CollectUsers  []User     `gorm:"many2many:article_collect_user" json:"collect_users"`
	LikeUsers     []User     `gorm:"many2many:article_like_user" json:"like_users"`
	Permission    uint8      `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence      uint8      `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at"`
	Status        uint64     `json:"status"`                        //状态
	ModifyTimes   uint64     `gorm:"default:0" json:"modify_times"` //修改次数
	ParentID      uint64     `json:"parent_id"`                     //父ID
	LastUser      User       `json:"last_user"`
	LastUserID    uint64     `json:"last_user_id"` //最后一个回复话题的人
	LastCommentAt *time.Time `json:"last_comment_at"`
}

const (
	// ArticleVerifying 审核中
	ArticleVerifying = iota

	// ArticleVerifySuccess 审核通过
	ArticleVerifySuccess

	// ArticleVerifyFail 审核未通过
	ArticleVerifyFail
)

const (
	//已删除
	ArticleDelete = iota

	ArticleNotDelete
)

// MaxTopArticleCount 最多能置顶的文章数
const MaxTopArticleCount = 4

const (
	// ContentTypeMarkdown markdown
	ContentTypeMarkdown = iota

	// ContentTypeHTML html
	ContentTypeHTML
)

type ArticleHistory struct {
	ID          uint64           `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time        `json:"created_at"`
	Title       string           `json:"title"`
	Content     string           `json:"content"`
	HTMLContent string           `json:"html_content"`
	ContentType int              `json:"content_type"`
	Categories  []Category       `gorm:"many2many:article_history_category" json:"categories"` //分类
	Tags        []Tag            `gorm:"many2many:article_history_tag" json:"tags"`
	User        User             `json:"user"`
	UserID      uint64           `json:"user_id"`
	Comments    []ArticleComment `gorm:"foreignkey:ArticleID" json:"comments"` //评论
	ActionCount
	ImageUrl    string `json:"image_url"`                              //图片
	ArticleID   uint64 `json:"article_id"`                             //根结点
	ParentID    uint64 `json:"parent_id"`                              //父节点
	ModifyTimes uint64 `json:"modify_times"`                           //修改次数
	DeleteFlag  uint8  `json:"delete_flag"`                            //是否删除
	Status      uint8  `gorm:"type:smallint ;default:0" json:"status"` //状态
}

type ArticleTag struct {
	ArticleID uint64 `gorm:"primary_key" json:"article_id"`
	TagName   string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}

type ArticleHistoryTag struct {
	ArticletHistoryID uint64 `gorm:"primary_key" json:"articlet_history_id"`
	TagName           string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}

type ArticleCategory struct {
	ArticleID  uint64 `gorm:"primary_key" json:"article_id"`
	CategoryID uint64 `gorm:"primary_key" json:"category_id"`
}
