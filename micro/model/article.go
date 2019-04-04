package model

import (
	"hoper/model/ov"
	"time"
)

type Article struct {
	ID          uint64              `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time           `json:"created_at"`
	Title       string              `gorm:"type:varchar(100)" json:"title"`
	Intro       string              `gorm:"type:varchar(100)" json:"intro"`
	Abstract    string              `gorm:"type:varchar(200)" json:"abstract"`
	Content     string              `gorm:"type:text" json:"content"`
	HTMLContent string              `gorm:"type:text" json:"html_content"`
	ContentType int                 `json:"content_type"`                                 //文本类型
	ImageUrl    string              `gorm:"type:varchar(100)" json:"image_url"`           //封面
	Categories  []ov.Category       `gorm:"many2many:article_category" json:"categories"` //分类
	SerialTitle string              `json:"serial_title"`
	Tags        []ov.Tag            `gorm:"many2many:article_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User        ov.User             `json:"user"`
	UserID      uint64              `json:"user_id"`
	Comments    []ov.ArticleComment `gorm:"ForeignKey:ArticleID" json:"comments"` //评论
	ActionCount
	ApproveUsers  []ov.User  `gorm:"many2many:article_approve" json:"approve_users"`
	CollectUsers  []ov.User  `gorm:"many2many:article_collect" json:"collect_users"`
	LikeUsers     []ov.User  `gorm:"many2many:article_like" json:"like_users"`
	Permission    uint8      `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence      uint8      `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `sql:"index" json:"-"`
	Status        uint8      `json:"-"`                  //状态
	ModifyTimes   uint8      `gorm:"default:0" json:"-"` //修改次数
	ParentID      uint64     `json:"-"`                  //修改的根节点
	LastUser      ov.User    `json:"last_user"`
	LastUserID    uint64     `json:"last_user_id"` //最后一个回复话题的人
	LastCommentAt *time.Time `json:"last_comment_at"`
}

type Serial struct {
	ID           uint64       `gorm:"primary_key" json:"id"`
	Title        string       `gorm:"type:varchar(100);unique_index" json:"title"`
	Description  string       `gorm:"type:varchar(500)" json:"description"` //描述
	Articles     []ov.Article `json:"articles"`
	ArticleCount uint64       `json:"article_count"`
	User         ov.User      `json:"user"`
	UserID       uint64       `json:"user_id"`
	ImageUrl     string       `gorm:"type:varchar(100)" json:"image_url"` //封面
	Count        uint64       `gorm:"default:0" json:"count"`
	Permission   uint8        `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence     uint8        `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
	DeletedAt    *time.Time   `sql:"index" json:"-"`
	Status       uint8        `gorm:"type:smallint;default:0" json:"-"` //状态
	ModifyTimes  uint8        `gorm:"default:0" json:"-"`               //修改次数
	ParentID     uint64       `json:"-"`                                //父节点
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

type ArticleTag struct {
	ArticleID uint64 `gorm:"primary_key" json:"article_id"`
	TagName   string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}

type ArticleCategory struct {
	ArticleID  uint64 `gorm:"primary_key" json:"article_id"`
	CategoryID uint64 `gorm:"primary_key" json:"category_id"`
}

type ArticleSerial struct {
	ArticleID uint64 `gorm:"primary_key" json:"article_id"`
	SerialID  uint64 `gorm:"primary_key" json:"serial_id"`
}

type ArticleApprove struct {
	CreatedAt time.Time  `json:"created_at"`
	ArticleID uint64     `gorm:"primary_key" json:"article_id"`
	UserID    uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type ArticleLike struct {
	CreatedAt time.Time  `json:"created_at"`
	ArticleID uint64     `gorm:"primary_key" json:"article_id"`
	UserID    uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

type ArticleCollect struct {
	CreatedAt time.Time  `json:"created_at"`
	ArticleID uint64     `gorm:"primary_key" json:"article_id"`
	UserID    uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}

//gorm:"many2many:PersonAccount;association_jointable_foreignkey:account_id;jointable_foreignkey:person_id"`
type ArticleBrowse struct {
	CreatedAt time.Time  `json:"created_at"`
	ArticleID uint64     `gorm:"primary_key" json:"article_id"`
	UserID    uint64     `gorm:"primary_key" json:"user_id"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"-"`
}
