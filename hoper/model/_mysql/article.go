package _mysql

import (
	"time"
)

type Article struct {
	ID            uint             `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time        `json:"created_at"`
	Title         string           `gorm:"type:varchar(100)" json:"title"`
	Content       string           `gorm:"type:longtext" json:"content"`
	HTMLContent   string           `gorm:"type:longtext" json:"html_content"`
	ContentType   int              `json:"content_type"`                                 //文本类型
	ImageUrl      string           `gorm:"type:varchar(100)" json:"image_url"`           //封面
	Categories    []Category       `gorm:"many2many:article_category" json:"categories"` //分类
	Tags          []Tag            `gorm:"many2many:article_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User          User             `json:"user"`
	UserID        uint             `json:"user_id"`
	Comments      []ArticleComment `json:"comments"`                       //评论
	BrowseCount   uint             `json:"browse_count"`                   //浏览
	CommentCount  uint             `gorm:"default:0" json:"comment_count"` //评论
	CollectCount  uint             `gorm:"default:0" json:"collect_count"` //收藏
	CollectUsers  []User           `gorm:"many2many:article_collection" json:"collect_users"`
	LoveCount     uint             `gorm:"default:0" json:"love_count"` //点赞
	LoveUsers     []User           `gorm:"many2many:article_love" json:"love_users"`
	Permission    uint8            `gorm:"type:tinyint(1) unsigned;default:0" json:"permission"` //查看权限
	DescFlag      uint8            `gorm:"type:tinyint(1) unsigned;default:0" json:"desc_flag"`  //排序，置顶
	UpdatedAt     *time.Time       `json:"updated_at"`
	DeletedAt     *time.Time       `sql:"index" json:"deleted_at"`
	Status        uint             `json:"status"`                        //状态
	ModifyTimes   uint             `gorm:"default:0" json:"modify_times"` //修改次数
	ParentID      uint             `json:"parent_id"`                     //父节点
	LastUser      User             `json:"last_user"`
	LastUserID    uint             `json:"last_user_id"` //最后一个回复话题的人
	LastCommentAt *time.Time       `json:"last_comment_at"`
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
	ID           uint             `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	Title        string           `json:"title"`
	Content      string           `json:"content"`
	HTMLContent  string           `json:"html_content"`
	ContentType  int              `json:"content_type"`
	Categories   []Category       `gorm:"many2many:article_history_category" json:"categories"` //分类
	Tags         []Tag            `gorm:"many2many:article_history_tag" json:"tags"`
	User         User             `json:"user"`
	UserID       uint             `json:"user_id"`
	Comments     []ArticleComment `gorm:"foreignkey:ArticleID" json:"comments"`             //评论
	BrowseCount  uint             `json:"browse_count"`                                     //浏览
	CommentCount uint             `json:"comment_count"`                                    //评论
	CollectCount uint             `json:"collect_count"`                                    //收藏
	LoveCount    uint             `json:"love_count"`                                       //点赞
	ImageUrl     string           `json:"image_url"`                                        //图片
	ArticleID    uint             `json:"article_id"`                                       //根结点
	ParentID     uint             `json:"parent_id"`                                        //父节点
	ModifyTimes  uint             `json:"modify_times"`                                     //修改次数
	DeleteFlag   uint             `json:"delete_flag"`                                      //是否删除
	Status       uint8            `gorm:"type:tinyint(1) unsigned;default:0" json:"status"` //状态
}
