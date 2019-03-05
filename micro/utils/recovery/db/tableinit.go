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
		&model.Career{},&model.Category{},&model.Collection{},&model.Like{},&model.CrawlerArticle{},
		&model.DiaryBookComment{},&model.MomentComment{},&model.FileUploadInfo{},&model.DiaryComment{},
		&model.Moment{},&model.Mood{},&model.DiaryBookHistory{},&model.DiaryHistory{},&model.ArticleHistory{},
			&model.MomentHistory{})*/

	/*

		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(
			&model.Tag{}, &model.Moment{},&model.Mood{},&model.MomentHistory{})*/

	//db.DropTable(&model.User{})
	/*	db.DropTableIfExists(&model.User{},
			&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.School{}, &model.Article{},
			&model.Career{}, &model.Category{}, &model.Collection{}, &model.Like{}, &model.CrawlerArticle{},
			&model.ArticleComment{},
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
		db.CreateTable(&model.User{},
			&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.School{}, &model.Article{},
			&model.Career{}, &model.Category{}, &model.Collection{}, &model.Like{}, &model.CrawlerArticle{},
			&model.ArticleComment{},
			&model.DiaryBookComment{}, &model.MomentComment{}, &model.FileUploadInfo{}, &model.DiaryComment{},
			&model.Moment{}, &model.Mood{}, &model.DiaryBookHistory{}, &model.DiaryHistory{}, &model.ArticleHistory{},
			&model.MomentHistory{})*/
	db.CreateTable(&model.School{})
	//db.CreateTable(&model.Career{})
}

/*
	SELECT concat('DROP TABLE IF EXISTS ', table_name, ';')
	FROM information_schema.tables
	WHERE table_schema = 'hoper';
*/
/*
清空所有表
CREATE OR REPLACE FUNCTION truncate_tables(username IN VARCHAR) RETURNS void AS $$
DECLARE
statements CURSOR FOR
SELECT tablename FROM pg_tables
WHERE tableowner = username AND schemaname = 'public';
BEGIN
FOR stmt IN statements LOOP
EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
END LOOP;
END;
$$ LANGUAGE plpgsql;
SELECT truncate_tables('postgres');
*/

/*
删除所有表
CREATE OR REPLACE FUNCTION del_tabs(username IN VARCHAR) RETURNS void AS $$
DECLARE
statements CURSOR FOR
SELECT tablename FROM pg_tables
WHERE tableowner = username AND schemaname = 'public';
BEGIN
FOR stmt IN statements LOOP
EXECUTE 'DROP TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
END LOOP;
END;
$$ LANGUAGE plpgsql;
SELECT del_tabs('postgres');*/
