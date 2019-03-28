package model

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
