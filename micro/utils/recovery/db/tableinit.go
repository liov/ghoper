package main

import (
	"hoper/initialize"
	"hoper/model"
)

func main() {

	db := initialize.DB
	/*
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{},
		&model.Tag{},&model.Diary{},&model.DiaryBook{},&model.School{},&model.Article{},
		&model.Career{},&model.Category{},&model.Collection{},&model.Love{},&model.CrawlerArticle{},
		&model.DiaryBookComment{},&model.MomentComment{},&model.FileUploadInfo{},&model.DiaryComment{},
		&model.Moment{},&model.Mood{},&model.DiaryBookHistory{},&model.DiaryHistory{},&model.ArticleHistory{},
			&model.MomentHistory{})*/

	/*

		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(
			&model.Tag{}, &model.Moment{},&model.Mood{},&model.MomentHistory{})*/

	//db.CreateTable(&model.User{})

	db.CreateTable(&model.User{},
		&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.School{}, &model.Article{},
		&model.Career{}, &model.Category{}, &model.Collection{}, &model.Love{}, &model.CrawlerArticle{},
		&model.DiaryBookComment{}, &model.MomentComment{}, &model.FileUploadInfo{}, &model.DiaryComment{},
		&model.Moment{}, &model.Mood{}, &model.DiaryBookHistory{}, &model.DiaryHistory{}, &model.ArticleHistory{},
		&model.MomentHistory{})
}

/*
	SELECT concat('DROP TABLE IF EXISTS ', table_name, ';')
	FROM information_schema.tables
	WHERE table_schema = 'hoper';
*/
