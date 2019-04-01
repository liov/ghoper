package model

import (
	"hoper/model/ov"
	"time"
)

type Moment struct {
	ID        uint64             `gorm:"primary_key" json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	Content   string             `gorm:"type:varchar(500)" json:"content"`
	ImageUrl  string             `json:"image_url"` //图片
	Mood      ov.Mood            `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName  string             `gorm:"type:varchar(20)" json:"mood_name"`
	Tags      []ov.Tag           `gorm:"many2many:moment_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments  []ov.MomentComment `json:"comments"` //评论
	User      ov.User            `json:"user"`
	UserID    uint64             `json:"user_id"`
	ActionCount
	ApproveUsers []ov.User  `gorm:"many2many:moment_approve_user" json:"approve_users"`
	CollectUsers []ov.User  `gorm:"many2many:moment_collect_user" json:"collect_users"`
	LikeUsers    []ov.User  `gorm:"many2many:moment_like_user" json:"like_users"`
	Sequence     uint8      `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Permission   uint8      `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Status       uint8      `gorm:"type:smallint;default:0" json:"-"`          //状态
	UpdatedAt    *time.Time `json:"-"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
	ModifyTimes  uint8      `gorm:"default:0" json:"-"` //修改次数
	ParentID     uint64     `json:"-"`                  //父节点
}

type MomentTag struct {
	MomentID uint64 `gorm:"primary_key" json:"momoent_id"`
	TagName  string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}
