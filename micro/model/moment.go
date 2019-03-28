package model

import (
	"time"
)

type Moment struct {
	ID              uint64          `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time       `json:"created_at"`
	Content         string          `gorm:"type:varchar(500)" json:"content"`
	ImageUrl        string          `gorm:"type:varchar(500)" json:"image_url"` //图片
	Mood            Mood            `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName        string          `gorm:"type:varchar(20)" json:"mood_name"`
	Tags            []Tag           `gorm:"many2many:moment_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments        []MomentComment `json:"comments"` //评论
	User            User            `json:"user"`
	UserID          uint64          `json:"user_id"`
	UpdatedAt       *time.Time      `json:"updated_at"`
	DeletedAt       *time.Time      `sql:"index" json:"deleted_at"`
	BrowseCount     uint64          `json:"browse_count"`                   //浏览
	CommentCount    uint64          `json:"comment_count"`                  //评论
	CollectCount    uint64          `json:"collect_count"`                  //收藏
	ApproveCount    uint64          `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers    []User          `gorm:"many2many:moment_approve_user" json:"approve_users"`
	CollectUsers    []User          `gorm:"many2many:moment_collect_user" json:"collect_users"`
	LikeCount       uint64          `gorm:"default:0" json:"like_count"` //喜欢
	LikeUsers       []User          `gorm:"many2many:moment_like_user" json:"like_users"`
	Sequence        uint8           `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Permission      uint8           `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Status          uint8           `gorm:"type:smallint;default:0" json:"status"`     //状态
	ModifyTimes     uint64          `gorm:"default:0" json:"modify_times"`             //修改次数
	MomentHistories []MomentHistory `json:"moment_histories"`
}

type MomentHistory struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	//EverCreatedAt time.Time       `json:"ever_created_at"`
	Content      string          `gorm:"type:varchar(500)" json:"content"`
	ImageUrl     string          `gorm:"type:varchar(100)" json:"image_url"` //图片
	Mood         Mood            `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName     string          `gorm:"type:varchar(20)" json:"mood_name"`
	Tags         []Tag           `gorm:"many2many:moment_history_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []MomentComment `gorm:"foreignkey:MomentID" json:"comments"` //评论
	User         User            `json:"user"`
	UserID       uint64          `json:"user_id"`
	BrowseCount  uint64          `json:"browse_count"`  //浏览
	CommentCount uint64          `json:"comment_count"` //评论
	CollectCount uint64          `json:"collect_count"` //收藏
	LikeCount    uint64          `json:"like_count"`    //点赞
	MomentID     uint64          `json:"moment_id"`     //根结点
	//ParentID     uint64            `json:"parent_id"`                                          //父节点
	ModifyTimes uint64 `gorm:"type:smallint" json:"modify_times"`          //修改次数
	DeleteFlag  uint8  `gorm:"type:smallint;default:0" json:"delete_flag"` //是否删除
	Status      uint8  `gorm:"type:smallint;default:0" json:"status"`      //状态
}

type MomentTag struct {
	MomentID uint64 `gorm:"primary_key" json:"momoent_id"`
	TagName  string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}

type MomentHistoryTag struct {
	MomentHistoryID uint64 `gorm:"primary_key" json:"moment_history_id"`
	TagName         string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}
