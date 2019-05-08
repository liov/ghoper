package ov

type KindOwn struct {
	Articles   []Article   `json:"articles"`
	Moments    []Moment    `json:"moments"`
	DiaryBooks []DiaryBook `json:"diary_books"`
	Diaries    []Diary     `json:"diaries"`
}

type KindOwnCount struct {
	MomentCount    uint64 `json:"moment_count"`
	ArticleCount   uint64 `json:"article_count"`
	DiaryBookCount uint64 `json:"diary_book_count"`
	DiaryCount     uint64 `json:"diary_count"`
	Count          uint64 `gorm:"default:0" json:"count"`
}
