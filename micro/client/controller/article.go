package controller

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/client/controller/common/e"
	"hoper/client/controller/common/gredis"
	"hoper/client/controller/common/logging"
	"hoper/initialize"
	"hoper/model"
	"hoper/utils"
	"time"

	"strconv"
	"strings"
	"unicode/utf8"
)

type Article struct {
	ID            uint             `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time        `json:"created_at"`
	Title         string           `gorm:"type:varchar(100)" json:"title"`
	Intro         string           `gorm:"type:varchar(100)" json:"intro"`
	Abstract      string           `gorm:"type:varchar(200)" json:"abstract"`
	Content       string           `gorm:"type:text" json:"content"`
	HTMLContent   string           `gorm:"type:text" json:"html_content"`
	ContentType   int              `json:"content_type"`                                 //文本类型
	ImageUrl      string           `gorm:"type:varchar(100)" json:"image_url"`           //封面
	Categories    []Category       `gorm:"many2many:article_category" json:"categories"` //分类
	Tags          []Tag            `gorm:"many2many:article_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User          User             `json:"user"`
	UserID        uint             `json:"user_id"`
	Comments      []ArticleComment `json:"comments"`                       //评论
	BrowseCount   uint             `json:"browse_count"`                   //浏览
	CommentCount  uint             `gorm:"default:0" json:"comment_count"` //评论
	CollectCount  uint             `gorm:"default:0" json:"collect_count"` //收藏
	CollectUsers  []User           `gorm:"many2many:article_collection" json:"collect_users"`
	LikeCount     uint             `gorm:"default:0" json:"like_count"` //点赞
	LikeUsers     []User           `gorm:"many2many:article_like" json:"like_users"`
	Permission    int8             `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Sort          int8             `gorm:"type:smallint;default:0" json:"sort"`       //排序，置顶
	UpdatedAt     *time.Time       `json:"updated_at"`
	DeletedAt     *time.Time       `sql:"index" json:"deleted_at"`
	Status        uint             `json:"status"`                        //状态
	ModifyTimes   uint             `gorm:"default:0" json:"modify_times"` //修改次数
	ParentID      uint             `json:"parent_id"`                     //父ID
	LastUser      User             `json:"last_user"`
	LastUserID    uint             `json:"last_user_id"` //最后一个回复话题的人
	LastCommentAt *time.Time       `json:"last_comment_at"`
}

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
			common.Response(c, articleCache)
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
	order := "sort desc," + orderStr + " desc"

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
		var tags []model.Tag
		var categories []model.Category
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

func articleValidation(c iris.Context, article *Article) (err error) {

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

	user := c.Values().Get("user").(User)

	if limitErr := common.Limit(model.ArticleMinuteLimit,
		model.ArticleMinuteLimitCount,
		model.ArticleDayLimit,
		model.ArticleMinuteLimitCount, user.ID); limitErr != "" {
		common.Response(c, limitErr, e.TimeTooMuch)
		return
	}

	var article Article

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
	article.ContentType = model.ContentTypeMarkdown
	article.ParentID = 0
	user.Score = user.Score + model.ArticleScore
	user.ArticleCount = user.ArticleCount + 1

	article.Title = utils.AvoidXSS(article.Title)
	article.Title = strings.TrimSpace(article.Title)

	if article.HTMLContent != "" {
		article.HTMLContent = utils.AvoidXSS(article.HTMLContent)
	}

	if s := []rune(article.Content); len(s) > 200 {

		article.Intro = string(s[:200])
		article.Abstract = string(s[:200])
	} else {
		article.Intro = article.Content
		article.Abstract = article.Content
	}

	saveErr := initialize.DB.Set("gorm:association_autocreate", false).Create(&article).Error
	nowTime := time.Now()
	for _, v := range article.Tags {
		if tag := ExistTagByName(v.Name); tag != nil {
			initialize.DB.Model(tag).Update("count", tag.Count+1)
		} else {
			newTag := Tag{CreatedAt: nowTime, Name: v.Name, Count: 1, UserID: user.ID}
			initialize.DB.Create(&newTag)
		}
		momentTag := model.ArticleTag{ArticleID: article.ID, TagName: v.Name}
		initialize.DB.Create(&momentTag)
	}

	for _, v := range article.Categories {
		articleCategory := model.ArticleCategory{ArticleID: article.ID, CategoryID: v.ID}
		initialize.DB.Create(&articleCategory)
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
		common.Response(c, "参数无效", e.ERROR)
		return
	}

	article.ParentID = historyArticle.ID
	article.ModifyTimes = article.ModifyTimes + 1

	saveErr := initialize.DB.Save(&article).Error

	if saveErr != nil {
		logging.Info("修改失败")
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
