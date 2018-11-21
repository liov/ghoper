package model

import "time"

// CrawlerArticle 爬虫抓取的文章
type CrawlerArticle struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	URL       string     `gorm:"type:varchar(100)" json:"url"`
	Title     string     `gorm:"type:varchar(100)" json:"title"`
	Content   string     `gorm:"type:text" json:"content"`
	From      int        `gorm:"type:varchar(100)" json:"from"`
	UpdatedAt *time.Time	`json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status        uint8  `gorm:"type:smallint;default:0" json:"status"`
}

const (
	// ArticleFromNULL 无来源
	ArticleFromNULL = 0

	// ArticleFromJianShu 简书
	ArticleFromJianShu = 1

	// ArticleFromZhihu 知乎
	ArticleFromZhihu = 2

	// ArticleFromHuxiu 虎嗅
	ArticleFromHuxiu = 3

	// ArticleFromCustom 自定义
	ArticleFromCustom = 10
)

const (
	// CrawlerScopePage 抓取单篇文章
	CrawlerScopePage = "page"

	// CrawlerScopeList 抓取一批文章
	CrawlerScopeList = "list"
)
