package credis

const (
	CacheArticle = "ARTICLE"
	CacheTag     = "TAG"
	CacheMoment  = "Moment"
	TopMoments   = "Moment_List_Top"
	Moments      = "Moment_List"
)

const (
	UserIndex = iota
	FlagIndex = iota + 10
	SysIndex
	CronIndex
)
