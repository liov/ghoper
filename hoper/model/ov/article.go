package ov

import "time"

type Article struct {
	ID            uint64     `gorm:"primary_key" json:"id"`
	Title         string     `gorm:"type:varchar(100)" json:"title"`
	Intro         string     `gorm:"type:varchar(100)" json:"intro"`
	Abstract      string     `gorm:"type:varchar(200)" json:"abstract"`
	Content       string     `gorm:"type:text" json:"content"`
	HTMLContent   string     `gorm:"type:text" json:"html_content"`
	ContentType   int        `json:"content_type"`                                 //文本类型
	ImageUrl      string     `gorm:"type:varchar(100)" json:"image_url"`           //封面
	Categories    []Category `gorm:"many2many:article_category" json:"categories"` //分类
	Tags          []Tag      `gorm:"many2many:article_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User          User       `json:"user"`
	UserID        uint64     `json:"user_id"`
	Permission    uint8      `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence      uint8      `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status        uint8      `json:"-"`
	ParentID      uint64     `json:"parent_id"` //父ID
	LastUser      User       `json:"last_user"`
	LastUserID    uint64     `json:"last_user_id"` //最后一个回复话题的人
	LastCommentAt *time.Time `json:"last_comment_at"`
}

type Serial struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	Title       string `gorm:"type:varchar(100);unique_index" json:"title"`
	Description string `gorm:"type:varchar(500)" json:"description"` //描述
}
