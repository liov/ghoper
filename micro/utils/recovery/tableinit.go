package main

import (
	"hoper/initialize"
	"hoper/model"
	"hoper/model/new/po"
)

//单独建个文件夹的目的很简单，测试的时候引用utils不会引用到这里
func main() {

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

	//initialize.DB.DropTable(&model.User{})

	/*	initialize.DB.Exec(`CREATE OR REPLACE FUNCTION del_tabs(username IN VARCHAR) RETURNS void AS $$
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
		initialize.DB.Exec(`SELECT del_tabs('postgres')`)*/
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
