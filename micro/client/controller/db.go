package controller

import (
	"github.com/kataras/iris"
	"hoper/initialize"
	"hoper/model"
)

func DBInit(c iris.Context) {
	initialize.DB.Exec(`CREATE OR REPLACE FUNCTION del_tabs(username IN VARCHAR) RETURNS void AS $$
DECLARE
statements CURSOR FOR
SELECT tablename FROM pg_tables
WHERE tableowner = username AND schemaname = 'public';
BEGIN
FOR stmt IN statements LOOP
EXECUTE 'DROP TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
END LOOP;
END;
$$ LANGUAGE plpgsql`)
	initialize.DB.Exec(`SELECT del_tabs('postgres')`)
	initialize.DB.CreateTable(&model.User{},
		&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.Education{}, &model.Article{},
		&model.Work{}, &model.Category{}, &model.Collection{}, &model.Like{}, &model.CrawlerArticle{},
		&model.ArticleComment{}, &model.Favorites{}, &model.Dictionary{}, &model.Follow{}, model.ArticleSet{},
		&model.DiaryBookComment{}, &model.MomentComment{}, &model.FileUploadInfo{}, &model.DiaryComment{},
		&model.Moment{}, &model.Mood{}, &model.DiaryBookHistory{}, &model.DiaryHistory{}, &model.ArticleHistory{},
		&model.MomentHistory{})
	initialize.DB.Create(&model.Category{Name: "未分类"})
	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()
	RedisConn.Do("flushall")
	c.WriteString("初始化成功")
}
