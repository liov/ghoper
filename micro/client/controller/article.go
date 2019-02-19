package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/client/controller/common/e"
	"hoper/client/controller/common/gredis"
	"hoper/client/controller/common/logging"
	"hoper/initialize"
	"hoper/model"
	"hoper/utils"

	"strconv"
	"strings"
	"unicode/utf8"
)

func GetArticle(c iris.Context) {

	var article, articleCache model.Article

	id := c.URLParam("id")

	key := "Article_" + id

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &articleCache)
			common.Response(c, e.SUCCESS, articleCache)
			return
		}
	}

	if err := initialize.DB.First(&article, id).Error; err != nil {
		common.Response(c, "无效的版块id")
		return
	}

	gredis.Set(key, article, 3600)
	common.Response(c, article)
}

func GetArticles(c iris.Context) {

	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))
	orderType := c.URLParam("orderType")

	tagID := c.URLParam("tagID")
	keyword := c.URLParam("keyword")
	categories := c.URLParam("categories")

	orderStr := "created_at"
	if orderType != "" {
		switch orderType {
		case "1":
			orderStr = "created_at"
		case "2":
			orderStr = "like_count"
		case "3":
			orderStr = "comment_count"
		}
	}
	order := "desc_flag desc" + orderStr + "desc"

	maps := map[string]interface{}{

		"tag_id": tagID,

		"keyword": keyword,

		"categories": categories,
	}

	var articles, cacheArticle []*model.Article

	key := strings.Join([]string{
		gredis.CacheArticle,
		"LIST",
		tagID, keyword, categories, orderStr,
	}, "_")

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			common.Response(c, e.SUCCESS, cacheArticle)
			return
		}
	}

	err := initialize.DB.Preload("Tag").Where(maps).Order(order).Offset(pageNo).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	gredis.Set(key, articles, 3600)

	common.Response(c, e.SUCCESS, articles)
}

func articleValidation(c iris.Context, article *model.Article) (err error) {

	err = &e.ValidtionError{Msg: "参数无效"}

	if article.Title == "" {
		common.Response(c, "文章名称不能为空")
		return
	}

	if utf8.RuneCountInString(article.Title) > model.MaxNameLen {
		msg := "文章名称不能超过" + strconv.Itoa(model.MaxNameLen) + "个字符"
		common.Response(c, msg)
		return
	}

	var theContent string

	if article.ContentType == model.ContentTypeHTML {
		theContent = article.HTMLContent
	} else {
		theContent = article.Content
	}

	if theContent == "" || utf8.RuneCountInString(theContent) <= 0 {
		common.Response(c, "文章内容不能为空")
		return
	}

	if utf8.RuneCountInString(theContent) > model.MaxContentLen {
		msg := "文章内容不能超过" + strconv.Itoa(model.MaxContentLen) + "个字符"
		common.Response(c, msg)
		return
	}

	if article.Tags == nil || len(article.Tags) <= 0 {
		common.Response(c, "请选择标签")
		return
	}

	if len(article.Categories) > model.MaxCategoryCount {
		msg := "文章最多属于" + strconv.Itoa(model.MaxCategoryCount) + "个版块"
		common.Response(c, msg)
		return
	}

	for i := 0; i < len(article.Categories); i++ {
		var category model.Category
		if err := initialize.DB.First(&category, article.Categories[i].ID).Error; err != nil {
			common.Response(c, "无效的版块id")
			return err
		}
		article.Categories[i] = category
	}

	return nil
}

// Create 创建文章
func AddArticle(c iris.Context) {

	user := c.GetViewData()["user"].(model.User)

	if limitErr := common.Limit(model.ArticleMinuteLimit,
		model.ArticleMinuteLimitCount,
		model.ArticleDayLimit,
		model.ArticleMinuteLimitCount, user.ID); limitErr != "" {
		common.Response(c, limitErr)
		return
	}

	var article model.Article

	if err := c.ReadJSON(&article); err != nil {
		fmt.Println(err.Error())
		common.Response(c, "参数无效")
		return
	}

	if err := articleValidation(c, &article); err != nil {
		common.Response(c, "参数无效")
		return
	}

	article.BrowseCount = 1
	article.Status = model.ArticleVerifying
	article.ModifyTimes = 0
	article.ContentType = model.ContentTypeMarkdown
	article.ParentID = 0
	user.Score = user.Score + model.ArticleScore
	user.ArticleCount = user.ArticleCount + 1
	/*if UserToRedis(user) != nil {
		common.SendErr(c,"error")
		return
	}*/

	article.Title = utils.AvoidXSS(article.Title)
	article.Title = strings.TrimSpace(article.Title)

	article.Content = strings.TrimSpace(article.Content)
	article.HTMLContent = strings.TrimSpace(article.HTMLContent)

	if article.HTMLContent != "" {
		article.HTMLContent = utils.AvoidXSS(article.HTMLContent)
	}

	saveErr := initialize.DB.Create(&article).Error

	if saveErr == nil {
		// 发表文章后，用户的积分、文章数会增加，如果保存失败了，不作处理
		if userErr := initialize.DB.Model(&user).Update(map[string]interface{}{
			"article_count": user.ArticleCount,
			"score":         user.Score,
		}).Error; userErr != nil {
			fmt.Println(userErr.Error())
		}
	}

	if saveErr != nil {
		return
	}

	common.Response(c, e.SUCCESS, "创建成功")
}

func historyArticle(c iris.Context, isDel uint) (model.ArticleHistory, model.Article, error) {

	var article model.Article
	//获取文章ID
	id := c.Params().GetUint64Default("id", 0)

	if err := initialize.DB.First(&article, id).Error; err != nil {
		common.Response(c, "无效的版块id")
		return model.ArticleHistory{}, model.Article{}, err
	}

	articleHistory := model.ArticleHistory{
		Title:       article.Title,
		ParentID:    article.ParentID,
		ArticleID:   article.ID,
		ModifyTimes: article.ModifyTimes + 1,
		DeleteFlag:  isDel,
		Content:     article.Content,
		HTMLContent: article.HTMLContent,
		ContentType: article.ContentType,
		Categories:  article.Categories,
		Tags:        article.Tags,
		Comments:    article.Comments,
		User:        article.User,
		UserID:      article.UserID,
		ImageUrl:    article.ImageUrl,
	}

	saveErr := initialize.DB.Create(&articleHistory).Error

	if saveErr != nil {
		logging.Info("保存历史失败")
	}

	return articleHistory, article, nil
}

//修改文章
func EditArticle(c iris.Context) {

	historyArticle, article, err := historyArticle(c, 0)

	if err != nil {
		return
	}

	if err := c.ReadJSON(&article); err != nil {
		common.Response(c, "参数无效")
		return
	}

	article.ParentID = historyArticle.ID
	article.ModifyTimes = article.ModifyTimes + 1

	saveErr := initialize.DB.Save(&article).Error

	if saveErr != nil {
		logging.Info("修改失败")
		return
	}

	common.Response(c, e.SUCCESS, "修改成功")
}

func DeleteArticle(c iris.Context) {

	historyArticle(c, 1)

	id := c.Params().GetUint64Default("id", 0)

	tx := initialize.DB.Begin()

	if err := tx.Where("id = ?", id).Delete(model.Article{}).Error; err != nil {
		common.Response(c, "error")
		tx.Rollback()
		return
	}
	common.Response(c, e.SUCCESS, "删除成功")
}
