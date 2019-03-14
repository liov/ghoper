package model

import (
	"time"
)

type Diary struct {
	ID           uint           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	Title        string         `gorm:"type:varchar(100)" json:"title"`
	DiaryBook    DiaryBook      `json:"diaryBook"`
	DiaryBookID  uint           `json:"diary_book_id"`
	Content      string         `json:"content"`
	Mood         Mood           `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName     string         `gorm:"type:varchar(20)" json:"mood_name"`
	Tags         []Tag          `gorm:"many2many:diary_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryComment `gorm:"ForeignKey:DiaryID" json:"comments"` //评论
	User         User           `json:"user"`
	UserID       uint           `json:"user_id"`
	ImageUrl     string         `json:"image_url"` //封面
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    *time.Time     `sql:"index" json:"deleted_at"`
	BrowseCount  uint           `json:"browse_count"`                   //浏览
	CommentCount uint           `gorm:"default:0" json:"comment_count"` //评论数
	CollectCount uint           `gorm:"default:0" json:"collect_count"` //收藏
	CollectUsers []User         `gorm:"-" json:"collect_users"`
	LikeCount    uint           `gorm:"default:0" json:"like_count"` //点赞
	LikeUsers    []User         `gorm:"many2many:diary_like" json:"like_users"`
	Permission   uint8          `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence     uint8          `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status       uint8          `gorm:"type:smallint;default:0" json:"status"`     //状态
	ModifyTimes  uint           `gorm:"default:0" json:"modify_times"`             //修改次数
	ParentID     uint           `json:"parent_id"`                                 //父节点
}

type DiaryHistory struct {
	ID           uint           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	Title        string         `gorm:"type:varchar(100)" json:"title"`
	DiaryBook    DiaryBook      `json:"diaryBook"`
	DiaryBookID  uint           `json:"diary_book_id"`
	Content      string         `json:"content"`
	Mood         Mood           `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName     string         `gorm:"type:varchar(20)" json:"mood_name"`
	Tags         []Tag          `gorm:"many2many:diary_history_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryComment `gorm:"ForeignKey:ArticleID" json:"comments"` //评论
	User         User           `json:"user"`
	UserID       uint           `json:"user_id"`
	ImageUrl     string         `json:"image_url"`     //封面
	BrowseCount  uint           `json:"browse_count"`  //浏览
	CommentCount uint           `json:"comment_count"` //评论数
	CollectCount uint           `json:"collect_count"` //收藏
	LikeCount    uint           `json:"like_count"`    //点赞
	DiaryID      uint           `json:"diary_id"`
	ParentID     uint           `json:"parent_id"`                                  //父节点
	ModifyTimes  uint           `json:"modify_times"`                               //修改次数
	DeleteFlag   uint8          `gorm:"type:smallint;default:0" json:"delete_flag"` //是否删除
	Status       uint8          `gorm:"type:smallint;default:0" json:"status"`      //状态
}

type DiaryBook struct {
	ID           uint               `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    *time.Time         `json:"updated_at"`
	DeletedAt    *time.Time         `sql:"index" json:"deleted_at"`
	Title        string             `gorm:"type:varchar(100)" json:"title"`
	Diaries      []Diary            `json:"diaries"`
	Description  string             `gorm:"type:varchar(500)" json:"description"` //描述
	Tags         []Tag              `gorm:"many2many:diary_book_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryBookComment `json:"comments"` //评论
	User         User               `json:"user"`
	UserID       uint               `json:"user_id"`
	ImageUrl     string             `gorm:"type:varchar(100)" json:"image_url"` //封面
	BrowseCount  uint               `json:"browse_count"`                       //浏览
	CommentCount uint               `gorm:"default:0" json:"comment_count"`     //评论数
	CollectCount uint               `gorm:"default:0" json:"collect_count"`     //收藏
	Collections  []Collection       `gorm:"many2many:diary_book_collection" json:"collections"`
	CollectUsers []User             `gorm:"-" json:"collect_users"`
	LikeCount    uint               `gorm:"default:0" json:"like_count"` //点赞
	LikeUsers    []User             `gorm:"many2many:diary_book_like" json:"like_users"`
	Permission   uint8              `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence     uint8              `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status       uint8              `gorm:"type:smallint;default:0" json:"status"`     //状态
	ModifyTimes  uint               `gorm:"default:0" json:"modify_times"`             //修改次数
	ParentID     uint               `json:"parent_id"`                                 //父节点
}

type DiaryBookHistory struct {
	ID           uint               `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time          `json:"created_at"`
	Title        string             `gorm:"type:varchar(100)" json:"title"`
	Diaries      []Diary            `json:"diaries"`
	Description  string             `gorm:"type:varchar(500)" json:"description"` //描述
	Tags         []Tag              `gorm:"many2many:diary_book_history_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryBookComment `json:"comments"` //评论
	User         User               `json:"user"`
	UserID       uint               `json:"user_id"`
	ImageUrl     string             `gorm:"type:varchar(100)" json:"image_url"` //封面
	BrowseCount  uint               `json:"browse_count"`                       //浏览
	CommentCount uint               `json:"comment_count"`                      //评论数
	CollectCount uint               `json:"collect_count"`                      //收藏
	LikeCount    uint               `json:"like_count"`                         //点赞
	DiaryBookID  uint               `json:"diary_book_id"`
	ParentID     uint               `json:"parent_id"`                                  //父节点
	ModifyTimes  uint               `json:"modify_times"`                               //修改次数
	DeleteFlag   uint8              `gorm:"type:smallint;default:0" json:"delete_flag"` //是否删除
	Status       uint8              `gorm:"type:smallint;default:0" json:"status"`      //状态
}
