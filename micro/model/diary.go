package model

import (
	"time"
)

type Diary struct {
	ID           uint64         `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	Title        string         `gorm:"type:varchar(100)" json:"title"`
	DiaryBook    DiaryBook      `json:"diaryBook"`
	DiaryBookID  uint64         `json:"diary_book_id"`
	Content      string         `json:"content"`
	Mood         Mood           `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName     string         `gorm:"type:varchar(20)" json:"mood_name"`
	Tags         []Tag          `gorm:"many2many:diary_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryComment `gorm:"ForeignKey:DiaryID" json:"comments"` //评论
	User         User           `json:"user"`
	UserID       uint64         `json:"user_id"`
	ImageUrl     string         `json:"image_url"` //封面
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    *time.Time     `sql:"index" json:"deleted_at"`
	BrowseCount  uint64         `json:"browse_count"`                   //浏览
	CommentCount uint64         `gorm:"default:0" json:"comment_count"` //评论数
	CollectCount uint64         `gorm:"default:0" json:"collect_count"` //收藏
	ApproveCount uint64         `gorm:"default:0" json:"approve_count"` //点赞
	ApproveUsers []User         `gorm:"many2many:diary_approve_user" json:"approve_users"`
	CollectUsers []User         `gorm:"many2many:diary_collect_user" json:"collect_users"`
	LikeCount    uint64         `gorm:"default:0" json:"like_count"` //喜欢
	LikeUsers    []User         `gorm:"many2many:diary_like_user" json:"like_users"`
	Permission   uint8          `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence     uint8          `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status       uint8          `gorm:"type:smallint;default:0" json:"status"`     //状态
	ModifyTimes  uint64         `gorm:"default:0" json:"modify_times"`             //修改次数
	ParentID     uint64         `json:"parent_id"`                                 //父节点
}

type DiaryHistory struct {
	ID           uint64         `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	Title        string         `gorm:"type:varchar(100)" json:"title"`
	DiaryBook    DiaryBook      `json:"diaryBook"`
	DiaryBookID  uint64         `json:"diary_book_id"`
	Content      string         `json:"content"`
	Mood         Mood           `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName     string         `gorm:"type:varchar(20)" json:"mood_name"`
	Tags         []Tag          `gorm:"many2many:diary_history_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryComment `gorm:"ForeignKey:ArticleID" json:"comments"` //评论
	User         User           `json:"user"`
	UserID       uint64         `json:"user_id"`
	ImageUrl     string         `json:"image_url"`     //封面
	BrowseCount  uint64         `json:"browse_count"`  //浏览
	CommentCount uint64         `json:"comment_count"` //评论数
	CollectCount uint64         `json:"collect_count"` //收藏
	LikeCount    uint64         `json:"like_count"`    //点赞
	DiaryID      uint64         `json:"diary_id"`
	ParentID     uint64         `json:"parent_id"`                                  //父节点
	ModifyTimes  uint64         `json:"modify_times"`                               //修改次数
	DeleteFlag   uint8          `gorm:"type:smallint;default:0" json:"delete_flag"` //是否删除
	Status       uint8          `gorm:"type:smallint;default:0" json:"status"`      //状态
}

type DiaryBook struct {
	ID           uint64             `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    *time.Time         `json:"updated_at"`
	DeletedAt    *time.Time         `sql:"index" json:"deleted_at"`
	Title        string             `gorm:"type:varchar(100)" json:"title"`
	Diaries      []Diary            `json:"diaries"`
	Description  string             `gorm:"type:varchar(500)" json:"description"` //描述
	Tags         []Tag              `gorm:"many2many:diary_book_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryBookComment `json:"comments"` //评论
	User         User               `json:"user"`
	UserID       uint64             `json:"user_id"`
	ImageUrl     string             `gorm:"type:varchar(100)" json:"image_url"` //封面
	BrowseCount  uint64             `json:"browse_count"`                       //浏览
	CommentCount uint64             `gorm:"default:0" json:"comment_count"`     //评论数
	CollectCount uint64             `gorm:"default:0" json:"collect_count"`     //收藏
	ApproveCount uint64             `gorm:"default:0" json:"approve_count"`     //点赞
	ApproveUsers []User             `gorm:"many2many:diary_book_approve_user" json:"approve_users"`
	CollectUsers []User             `gorm:"many2many:diary_book_collect_user" json:"collect_users"`
	LikeCount    uint64             `gorm:"default:0" json:"like_count"` //喜欢
	LikeUsers    []User             `gorm:"many2many:diary_book_like_user" json:"like_users"`
	Permission   uint8              `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sequence     uint8              `gorm:"type:smallint;default:0" json:"sequence"`   //排序，置顶
	Status       uint8              `gorm:"type:smallint;default:0" json:"status"`     //状态
	ModifyTimes  uint64             `gorm:"default:0" json:"modify_times"`             //修改次数
	ParentID     uint64             `json:"parent_id"`                                 //父节点
}

type DiaryBookHistory struct {
	ID           uint64             `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time          `json:"created_at"`
	Title        string             `gorm:"type:varchar(100)" json:"title"`
	Diaries      []Diary            `json:"diaries"`
	Description  string             `gorm:"type:varchar(500)" json:"description"` //描述
	Tags         []Tag              `gorm:"many2many:diary_book_history_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	Comments     []DiaryBookComment `json:"comments"` //评论
	User         User               `json:"user"`
	UserID       uint64             `json:"user_id"`
	ImageUrl     string             `gorm:"type:varchar(100)" json:"image_url"` //封面
	BrowseCount  uint64             `json:"browse_count"`                       //浏览
	CommentCount uint64             `json:"comment_count"`                      //评论数
	CollectCount uint64             `json:"collect_count"`                      //收藏
	LikeCount    uint64             `json:"like_count"`                         //点赞
	DiaryBookID  uint64             `json:"diary_book_id"`
	ParentID     uint64             `json:"parent_id"`                                  //父节点
	ModifyTimes  uint64             `json:"modify_times"`                               //修改次数
	DeleteFlag   uint8              `gorm:"type:smallint;default:0" json:"delete_flag"` //是否删除
	Status       uint8              `gorm:"type:smallint;default:0" json:"status"`      //状态
}
