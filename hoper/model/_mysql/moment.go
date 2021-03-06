package _mysql

import (
	"time"
)

type Moment struct {
	ID             uint            `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time       `json:"created_at"`
	Content        string          `gorm:"type:varchar(500)" json:"content"`
	ImageUrl       string          `gorm:"type:varchar(100)" json:"image_url"` //图片
	Mood           Mood            `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName       string          `gorm:"type:varchar(20)" json:"mood_name"`
	Tags           []Tag           `gorm:"many2many:moment_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments       []MomentComment `json:"comments"` //评论
	User           User            `json:"user"`
	UserID         uint            `json:"user_id"`
	UpdatedAt      *time.Time      `json:"updated_at"`
	DeletedAt      *time.Time      `sql:"index" json:"deleted_at"`
	BrowseCount    uint            `json:"browse_count"`  //浏览
	CommentCount   uint            `json:"comment_count"` //评论
	CollectCount   uint            `json:"collect_count"` //收藏
	CollectUsers   []User          `gorm:"many2many:moment_collection" json:"collect_users"`
	LoveCount      uint            `json:"love_count"` //点赞
	LoveUsers      []User          `gorm:"many2many:moment_love" json:"love_users"`
	DescFlag       uint8           `gorm:"type:tinyint unsigned;default:0" json:"desc_flag"`     //排序，置顶
	Permission     uint8           `gorm:"type:tinyint(1) unsigned;default:0" json:"permission"` //查看权限
	Status         uint8           `gorm:"type:tinyint(1) unsigned;default:0" json:"status"`     //状态
	ModifyTimes    uint            `gorm:"default:0" json:"modify_times"`                        //修改次数
	MomentHistorys []MomentHistory `json:"moment_historys"`
}

type MomentHistory struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	//EverCreatedAt time.Time       `json:"ever_created_at"`
	Content      string          `gorm:"type:varchar(500)" json:"content"`
	ImageUrl     string          `gorm:"type:varchar(100)" json:"image_url"` //图片
	Mood         Mood            `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName     string          `gorm:"type:varchar(20)" json:"mood_name"`
	Tags         []Tag           `gorm:"many2many:moment_history_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []MomentComment `gorm:"foreignkey:MomentID" json:"comments"` //评论
	User         User            `json:"user"`
	UserID       uint            `json:"user_id"`
	BrowseCount  uint            `json:"browse_count"`  //浏览
	CommentCount uint            `json:"comment_count"` //评论
	CollectCount uint            `json:"collect_count"` //收藏
	LoveCount    uint            `json:"love_count"`    //点赞
	MomentID     uint            `json:"moment_id"`     //根结点
	//ParentID     uint            `json:"parent_id"`                                          //父节点
	ModifyTimes uint  `gorm:"type:tinyint unsigned" json:"modify_times"`          //修改次数
	DeleteFlag  uint8 `gorm:"type:tinyint unsigned;default:0" json:"delete_flag"` //是否删除
	Status      uint8 `gorm:"type:tinyint(1) unsigned;default:0" json:"status"`   //状态
}

type MomentTag struct {
	MomentID uint   `gorm:"primary_key" json:"momoent_id"`
	TagName  string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}

type MomentHistoryTag struct {
	MomentHistoryID uint   `gorm:"primary_key" json:"moment_history_id"`
	TagName         string `gorm:"type:varchar(10);primary_key" json:"tag_name"`
}
