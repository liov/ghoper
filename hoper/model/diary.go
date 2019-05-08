package model

import (
	"time"
)

type Diary struct {
	ID          uint64         `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	Title       string         `gorm:"type:varchar(100)" json:"title"`
	DiaryBook   DiaryBook      `json:"diaryBook"`
	DiaryBookID uint64         `json:"diary_book_id"`
	Content     string         `json:"content"`
	Mood        Mood           `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName    string         `gorm:"type:varchar(20)" json:"mood_name"`
	Tags        []Tag          `gorm:"many2many:diary_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments    []DiaryComment `gorm:"ForeignKey:DiaryID" json:"comments"` //评论
	User        User           `json:"user"`
	UserID      uint64         `json:"user_id"`
	ImageUrl    string         `json:"image_url"` //封面
	ActionCount
	ApproveUsers []User     `gorm:"many2many:diary_approve_user" json:"approve_users"`
	CollectUsers []User     `gorm:"many2many:diary_collect_user" json:"collect_users"`
	LikeUsers    []User     `gorm:"many2many:diary_like_user" json:"like_users"`
	Permission   uint8      `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence     uint8      `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status       uint8      `gorm:"type:smallint;default:0" json:"-"`          //状态
	UpdatedAt    *time.Time `json:"-"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
	ModifyTimes  uint8      `gorm:"default:0" json:"-"` //修改次数
	ParentID     uint64     `json:"-"`                  //父节点
}

type DiaryBook struct {
	ID          uint64             `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	Title       string             `gorm:"type:varchar(100)" json:"title"`
	Diaries     []Diary            `json:"diaries"`
	Description string             `gorm:"type:varchar(500)" json:"description"` //描述
	Tags        []Tag              `gorm:"many2many:diary_book_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments    []DiaryBookComment `json:"comments"` //评论
	User        User               `json:"user"`
	UserID      uint64             `json:"user_id"`
	ImageUrl    string             `gorm:"type:varchar(100)" json:"image_url"` //封面
	ActionCount
	ApproveUsers []User     `gorm:"many2many:diary_book_approve_user" json:"approve_users"`
	CollectUsers []User     `gorm:"many2many:diary_book_collect_user" json:"collect_users"`
	LikeUsers    []User     `gorm:"many2many:diary_book_like_user" json:"like_users"`
	Permission   uint8      `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence     uint8      `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status       uint8      `gorm:"type:smallint;default:0" json:"-"`          //状态
	UpdatedAt    *time.Time `json:"-"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
	ModifyTimes  uint8      `gorm:"default:0" json:"-"` //修改次数
	ParentID     uint64     `json:"parent_id"`          //父节点
}
