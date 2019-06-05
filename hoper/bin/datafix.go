package main

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"hoper/client/controller"
	"hoper/initialize"
	"hoper/model"
	"hoper/utils"
	"hoper/utils/ulog"
	"os"
)

func main() {

	if log, ok := ulog.Log.(*zap.SugaredLogger); ok {
		defer log.Sync()
	}

	defer initialize.DB.Close()

	Article()
}

func Article() {
	type DataFix struct {
		Data []model.Article `json:"data"`
	}

	f, err := os.Open("../../static/datafix/copy.json")
	if err != nil {
		ulog.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			ulog.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	s.Buffer(nil, 200*1024)
	var dataFix DataFix
	for s.Scan() {
		utils.Json.Unmarshal(s.Bytes(), &dataFix)
		for i := len(dataFix.Data) - 1; i >= 0; i-- {
			fmt.Println(dataFix.Data[i].ID)
			dataFix.Data[i].ID = 0
			initialize.DB.Set("gorm:association_autocreate", false).Create(&dataFix.Data[i])

			for _, v := range dataFix.Data[i].Tags {
				articleTag := model.ArticleTag{ArticleID: dataFix.Data[i].ID, TagName: v.Name}
				initialize.DB.Create(&articleTag)
			}

			for _, v := range dataFix.Data[i].Categories {
				articleCategory := model.ArticleCategory{ArticleID: dataFix.Data[i].ID, CategoryID: v.ID}
				initialize.DB.Create(&articleCategory)

			}

			if serialID := controller.CreatSerial(&dataFix.Data[i].SerialTitle, 1); serialID > 0 {
				articleSerial := model.ArticleSerial{ArticleID: dataFix.Data[i].ID, SerialID: serialID}
				initialize.DB.Create(&articleSerial)

			}

		}
	}
	err = s.Err()
	if err != nil {
		ulog.Fatal(err)
	}
}
