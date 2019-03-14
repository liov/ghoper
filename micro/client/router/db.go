package router

import (
	"github.com/kataras/iris"
	"hoper/initialize"
	"hoper/model"
)

func DBInit(c iris.Context) {
	initialize.DB.DropTableIfExists(&model.User{},
		&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.Education{}, &model.Article{},
		&model.Work{}, &model.Category{}, &model.Collection{}, &model.Like{}, &model.CrawlerArticle{},
		&model.ArticleComment{}, &model.Favorites{},
		&model.DiaryBookComment{}, &model.MomentComment{}, &model.FileUploadInfo{}, &model.DiaryComment{},
		&model.Moment{}, &model.Mood{}, &model.DiaryBookHistory{}, &model.DiaryHistory{}, &model.ArticleHistory{},
		&model.MomentHistory{}, "article_category", "article_collection", "article_comment", "article_history",
		"article_history_category", "article_history_tag", "article_like", "article_tag",
		"diary_book_category", "diary_book_collection", "diary_book_comment", "diary_book_history",
		"diary_book_history_category", "diary_book_history_tag", "diary_book_like", "diary_book_tag",
		"diary_category", "diary_collection", "diary_comment", "diary_history",
		"diary_history_category", "diary_history_tag", "diary_like", "diary_tag",
		"moment_category", "moment_collection", "moment_comment", "moment_history",
		"moment_history_category", "moment_history_tag", "moment_like", "moment_tag", "user_collection")
	initialize.DB.CreateTable(&model.User{},
		&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.Education{}, &model.Article{},
		&model.Work{}, &model.Category{}, &model.Collection{}, &model.Like{}, &model.CrawlerArticle{},
		&model.ArticleComment{}, &model.Favorites{},
		&model.DiaryBookComment{}, &model.MomentComment{}, &model.FileUploadInfo{}, &model.DiaryComment{},
		&model.Moment{}, &model.Mood{}, &model.DiaryBookHistory{}, &model.DiaryHistory{}, &model.ArticleHistory{},
		&model.MomentHistory{})
	initialize.DB.Create(&model.Category{Name: "未分类"})
	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()
	RedisConn.Do("flushall")
	c.WriteString("初始化成功")
}
