package vo

import "time"

type CUDTime struct {
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"` //状态
}

type KindOwn struct {
	Articles []Article `json:"articles"`
	Moments  []Moment  `json:"moments"`
	//DiaryBooks []DiaryBook `json:"diary_books"`
	//Diaries    []Diary     `json:"diaries"`
}

type ActionCount struct {
	CollectCount int64 `gorm:"default:0" json:"collect_count"` //收藏
	LikeCount    int64 `gorm:"default:0" json:"like_count"`    //喜欢
	ApproveCount int64 `gorm:"default:0" json:"approve_count"` //点赞
	CommentCount int64 `gorm:"default:0" json:"comment_count"` //评论
	BrowseCount  int64 `gorm:"default:0" json:"browse_count"`  //浏览
}

type KindOwnCount struct {
	MomentCount    uint64 `json:"moment_count"`
	ArticleCount   uint64 `json:"article_count"`
	DiaryBookCount uint64 `json:"diary_book_count"`
	DiaryCount     uint64 `json:"diary_count"`
	Count          uint64 `gorm:"default:0" json:"count"`
}
