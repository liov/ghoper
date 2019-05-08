package mysqlmodel

// 通用常量

const (
	// DevelopmentMode 开发模式
	DevelopmentMode = "development"

	// TestMode 测试模式
	TestMode = "hoper"

	// ProductionMode 产品模式
	ProductionMode = "production"
)

const (
	// NoParent 无父结点时的parent_id
	NoParent = 0

	// MaxOrder 最大的排序号
	MaxOrder = 10000

	// MinOrder 最小的排序号
	MinOrder = 0

	// PageSize 默认每页的条数
	PageSize = 5

	// MaxPageSize 每页最大的条数
	MaxPageSize = 100

	// MinPageSize 每页最小的条数
	MinPageSize = 5

	// MaxNameLen 最大的名称长度
	MaxNameLen = 100

	// MaxContentLen 最大的内容长度
	MaxContentLen = 50000

	// MaxCategoryCount 最多可以属于几个分类
	MaxCategoryCount = 6
)

const (

	// ArticleMinuteLimitCount 用户每分钟最多能发表的文章数
	MomentMinuteLimitCount = 30

	// ArticleDayLimitCount 用户每天最多能发表的文章数
	MomentDayLimitCount = 1000

	// ArticleMinuteLimitCount 用户每分钟最多能发表的文章数
	ArticleMinuteLimitCount = 30

	// ArticleDayLimitCount 用户每天最多能发表的文章数
	ArticleDayLimitCount = 1000

	// CommentMinuteLimitCount 用户每分钟最多能发表的评论数
	CommentMinuteLimitCount = 30

	// CommentDayLimitCount 用户每天最多能发表的评论数
	CommentDayLimitCount = 1000
)
