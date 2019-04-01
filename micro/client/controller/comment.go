package controller

import (
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"

	"time"
)

type ArticleComment struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint64    `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	ArticleID uint64    `json:"article_id"` //话题或投票的ID
	ParentID  uint64    `json:"parent_id"`  //直接父评论的ID
}

type MomentComment struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint64    `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	MomentID  uint64    `json:"moment_id"` //瞬间ID
	ParentID  uint64    `json:"parent_id"` //直接父评论的ID
}

type DiaryComment struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint64    `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	DiaryID   uint64    `json:"diary_id"`  //瞬间ID
	ParentID  uint64    `json:"parent_id"` //直接父评论的ID
}

type DiaryBookComment struct {
	ID          uint64    `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	UserID      uint64    `json:"user_id"`
	Content     string    `gorm:"type:varchar(500)" json:"content"`
	HTMLContent string    `gorm:"type:varchar(500)" json:"html_content"`
	ContentType int       `json:"content_type"`
	DiaryBookID uint64    `json:"diary_book_id"` //瞬间ID
	ParentID    uint64    `json:"parent_id"`     //直接父评论的ID
}

type Comment interface {
	GetCommentType()
}

func AddComment(c iris.Context) {
	userID := c.Values().Get("userID").(uint64)
	if limitErr := common.Limit(model.CommentMinuteLimit,
		model.CommentMinuteLimitCount,
		model.CommentDayLimit,
		model.CommentMinuteLimitCount, userID); limitErr != nil {
		common.Response(c, limitErr.Error())
		return
	}
	kind := c.Params().Get("kind")
	nowTime := time.Now()
	switch KindIndex[kind] {
	case kindArticle:
		var comment ArticleComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case kindMoment:
		var comment MomentComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case kindDiary:
		var comment DiaryComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case kindDiaryBook:
		var comment DiaryBookComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	}

	common.Response(c, "评论成功")
}

func commentBind(comment interface{}, c iris.Context) {
	err := c.ReadJSON(comment)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"controller": "moment",
		}).Info(err.Error())
		common.Response(c, "参数无效")
		return
	}
}

func GetComment(ctx iris.Context) {
	userID := ctx.Values().Get("userID").(uint64)
	kind := ctx.Params().Get("kind")
	nowTime := time.Now()
	switch KindIndex[kind] {
	case kindArticle:
		var comment ArticleComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case kindMoment:
		var comment MomentComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case kindDiary:
		var comment DiaryComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case kindDiaryBook:
		var comment DiaryBookComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	}
}
