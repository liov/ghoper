package model

import "time"

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/3/28
 * @description：
 */

type ActionCount struct {
	CollectCount int64 `gorm:"default:0" json:"collect_count"` //收藏
	LikeCount    int64 `gorm:"default:0" json:"like_count"`    //喜欢
	ApproveCount int64 `gorm:"default:0" json:"approve_count"` //点赞
	CommentCount int64 `gorm:"default:0" json:"comment_count"` //评论
	BrowseCount  int64 `gorm:"default:0" json:"browse_count"`  //浏览
}

type Favorites struct {
	ID          uint64       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	Name        string       `gorm:"type:varchar(20)" json:"name"`
	User        User         `json:"user"`
	UserID      uint64       `json:"user_id"`
	FollowUsers []User       `json:"follow_users"`
	Count       uint64       `json:"count"`
	Collections []Collection `gorm:"many2many:collection_favorites" json:"collections"`
	UpdatedAt   *time.Time   `json:"updated_at"`
	DeletedAt   *time.Time   `sql:"index" json:"deleted_at"`
	Status      uint8        `gorm:"type:smallint;default:0" json:"status"`
}

//收藏夹？像网易云一样可以收藏别人的歌单
type Collection struct {
	ID          uint64      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	RefID       uint64      `json:"ref_id"`
	Kind        string      `gorm:"type:varchar(10)" json:"kind"`
	Favorites   []Favorites `json:"favorites"`
	FavoritesID uint64      `json:"favorites_id"`
	UserID      uint64      `json:"user_id"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `sql:"index" json:"deleted_at"`
	Status      uint8       `gorm:"type:smallint;default:0" json:"status"`
}

type Like struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	RefID     uint64     `json:"ref_id"`
	Kind      string     `gorm:"type:varchar(10)" json:"kind"`
	User      User       `json:"user"`
	UserID    uint64     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

type ArticleCollection struct {
	ArticleID   uint64       `json:"article_id"`
	Collections []Collection `json:"collections"`
}
