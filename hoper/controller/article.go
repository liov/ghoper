package controller

import (
	"hoper/utils/ulog"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"hoper/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/crm"
	"hoper/model/e"
	"hoper/model/ov"
	"hoper/utils"
)

func GetArticle(c iris.Context) {

	var article model.Article

	id := c.Params().GetUint64Default("id", 0)

	if err := initialize.DB.First(&article, id).Error; err != nil {
		common.Response(c, "无效的文章id")
		return
	}
	var tags []ov.Tag
	var categories []ov.Category
	initialize.DB.Model(&article).Related(&tags, "Tags").Related(&categories, "Categories")
	article.Tags = tags
	article.Categories = categories

	//gredis.Set(key, article, 3600)
	common.Response(c, article)
}

func GetArticles(c iris.Context) {

	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))
	orderType := c.URLParam("orderType")

	/*	tagID := c.URLParam("tagID")
		keyword := c.URLParam("keyword")
		categories := c.URLParam("categories")*/

	orderStr := "created_at"

	if orderType != "" {
		orderStr = orderType
	}
	/*		switch orderType {
			case "1":
				orderStr = "created_at"
			case "2":
				orderStr = "like_count"
			case "3":
				orderStr = "comment_count"
			}
		}*/
	order := "sequence desc," + orderStr + " desc"

	/*	maps := map[string]interface{}{

		"tag_id": tagID,

		"categories": categories,
	}*/

	var articles []model.Article

	/*var  cacheArticles []model.Article
	key := strings.Join([]string{
		gredis.CacheArticle,
		"LIST",
		tagID, keyword, categories, orderStr,
	}, "_")

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if gredis.Exists(key) {
		data, err := redis.Bytes(conn.Do("GET", key))
		count, err := redis.Int(conn.Do("GET", key+"_count"))
		utils.Json.Unmarshal(data, &cacheArticles)
		if err != nil {
			logging.Info(err)
		} else {
			common.Res(c, iris.Map{
				"data":  cacheArticles,
				"count": count,
				"msg":   e.GetMsg(e.SUCCESS),
				"code":  e.SUCCESS,
			})
			return
		}
	}*/
	var count int
	err := initialize.DB.Preload("User").Order(order).Offset(pageNo * pageSize).Limit(pageSize).Find(&articles).Error
	initialize.DB.Table("article").Count(&count)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	for i, a := range articles {
		var tags []ov.Tag
		var categories []ov.Category
		initialize.DB.Model(&a).Related(&tags, "Tags").Related(&categories, "Categories")
		articles[i].Tags = tags
		articles[i].Categories = categories
	}

	/*	as, _ := utils.Json.Marshal(articles)
		conn.Do("SET", key, as)
		conn.Do("EXPIRE", key, 3600)
		conn.Do("SET", key+"_count", strconv.Itoa(count))
		conn.Do("EXPIRE", key+"_count", 3600)*/

	common.Res(c, iris.Map{
		"data":  articles,
		"count": count,
		"msg":   e.GetMsg(e.SUCCESS),
		"code":  e.SUCCESS,
	})
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

	if theContent == "\n" || utf8.RuneCountInString(theContent) < 20 {
		common.Response(c, "文章内容不能小于20个字")
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
		common.Response(c, msg, e.ERROR)
		return
	}

	/*	for i := 0; i < len(article.Categories); i++ {
			var category model.Category
			if err := initialize.DB.First(&category, article.Categories[i]).Error; err != nil {
				common.Response(c, "无效的版块id")
				return err
			}
			article.Categories[i] = category.ID
		}
	*/
	return nil
}

// Create 创建文章
func AddArticle(c iris.Context) {

	user := c.Values().Get("user").(*User)

	if limitErr := common.Limit(model.ArticleMinuteLimit,
		model.ArticleMinuteLimitCount,
		model.ArticleDayLimit,
		model.ArticleMinuteLimitCount, user.ID); limitErr != nil {
		common.Response(c, limitErr.Error(), e.TimeTooMuch)
		return
	}

	var article model.Article

	if err := c.ReadJSON(&article); err != nil {
		common.Response(c, "参数无效", e.ERROR)
		return
	}

	if err := articleValidation(c, &article); err != nil {
		return
	}

	article.UserID = user.ID
	article.BrowseCount = 1
	article.Status = model.ArticleVerifySuccess
	article.ModifyTimes = 0
	article.ParentID = 0
	user.Score = user.Score + model.ArticleScore
	user.ArticleCount = user.ArticleCount + 1

	article.Title = utils.AvoidXSS(article.Title)
	article.Title = strings.TrimSpace(article.Title)

	if article.HTMLContent != "" {
		article.HTMLContent = utils.AvoidXSS(article.HTMLContent)
	}

	if s := []rune(article.Content); len(s) > 200 {

		article.Intro = string(s[:100])
		article.Abstract = string(s[:200])
	} else {
		article.Intro = article.Content
		article.Abstract = article.Content
	}

	article.CreatedAt = time.Now()
	article.UpdatedAt = &article.CreatedAt
	saveErr := initialize.DB.Set("gorm:association_autocreate", false).Create(&article).Error

	for _, v := range article.Tags {
		if ExistTagByName(&v, user.ID) {
			setFlagCountToRedis(flagTag, v.Name, 1)
		}
		articleTag := model.ArticleTag{ArticleID: article.ID, TagName: v.Name}
		initialize.DB.Create(&articleTag)
	}

	for _, v := range article.Categories {
		articleCategory := model.ArticleCategory{ArticleID: article.ID, CategoryID: v.ID}
		initialize.DB.Create(&articleCategory)
		setFlagCountToRedis(flagCategory, v.Name, 1)
	}

	if serialID := CreatSerial(&article.SerialTitle, user.ID); serialID > 0 {
		articleSerial := model.ArticleSerial{ArticleID: article.ID, SerialID: serialID}
		initialize.DB.Create(&articleSerial)

	}

	if saveErr == nil {
		// 发表文章后，用户的积分、文章数会增加，如果保存失败了，不作处理
		if userErr := initialize.DB.Model(&user).Update(map[string]interface{}{
			"article_count": user.ArticleCount,
			"score":         user.Score,
		}).Error; userErr != nil {

		}
	}

	if saveErr != nil {
		return
	}

	common.Response(c, "保存成功", e.SUCCESS)
}

func historyArticle(c iris.Context, isDel uint8) (*crm.ArticleHistory, *model.Article, error) {

	var article model.Article
	//获取文章ID
	id := c.Params().GetUint64Default("id", 0)

	if err := initialize.DB.First(&article, id).Error; err != nil {
		common.Response(c, "无效的版块id")
		return nil, nil, err
	}

	articleHistory := crm.ArticleHistory{
		Title:       article.Title,
		CreatedAt:   time.Now(),
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
		UserID:      article.UserID,
		ImageUrl:    article.ImageUrl,
	}

	saveErr := initialize.DB.Create(&articleHistory).Error

	if saveErr != nil {
		ulog.Info("保存历史失败")
	}

	return &articleHistory, &article, nil
}

// 修改文章
func EditArticle(c iris.Context) {

	historyArticle, article, err := historyArticle(c, 0)

	if err != nil {
		return
	}

	if err := c.ReadJSON(article); err != nil {
		common.Response(c, "参数无效", e.ERROR)
		return
	}

	if err := articleValidation(c, article); err != nil {
		return
	}

	article.Title = utils.AvoidXSS(article.Title)
	article.Title = strings.TrimSpace(article.Title)

	if article.HTMLContent != "" {
		article.HTMLContent = utils.AvoidXSS(article.HTMLContent)
	}

	if s := []rune(article.Content); len(s) > 200 {

		article.Intro = string(s[:100])
		article.Abstract = string(s[:200])
	} else {
		article.Intro = article.Content
		article.Abstract = article.Content
	}

	article.ParentID = historyArticle.ID
	article.ModifyTimes = article.ModifyTimes + 1

	saveErr := initialize.DB.Set("gorm:save_associations", false).Save(article).Error

	if saveErr != nil {
		common.Response(c, "修改成功", e.ERROR)
		return
	}

	common.Response(c, "修改成功", e.SUCCESS)
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
