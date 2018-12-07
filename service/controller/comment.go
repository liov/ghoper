package controller

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"service/controller/common"
	"service/initialize"
	"service/model"
	"time"
)

type ArticleComment struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint      `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	ArticleID uint      `json:"article_id"` //话题或投票的ID
	ParentID  uint      `json:"parent_id"`  //直接父评论的ID
}

type MomentComment struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint      `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	MomentID  uint      `json:"moment_id"` //瞬间ID
	ParentID  uint      `json:"parent_id"` //直接父评论的ID
}

type DiaryComment struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	UserID    uint      `json:"user_id"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	DiaryID   uint      `json:"diary_id"`  //瞬间ID
	ParentID  uint      `json:"parent_id"` //直接父评论的ID
}

type DiaryBookComment struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	UserID      uint      `json:"user_id"`
	Content     string    `gorm:"type:varchar(500)" json:"content"`
	HTMLContent string    `gorm:"type:varchar(500)" json:"html_content"`
	ContentType int       `json:"content_type"`
	DiaryBookID uint      `json:"diary_book_id"` //瞬间ID
	ParentID    uint      `json:"parent_id"`     //直接父评论的ID
}

type Comment interface {
	GetCommentType()
}

func AddComment(c *fasthttp.RequestCtx) {
	user := c.UserValue("user").(User)
	if limitErr := common.Limit(model.CommentMinuteLimit,
		model.CommentMinuteLimitCount,
		model.CommentDayLimit,
		model.CommentMinuteLimitCount, user.ID); limitErr != "" {
		common.Response(c, limitErr)
		return
	}
	classify := c.UserValue("classify").(string)
	nowTime := time.Now()
	switch classify {
	case "articleComment":
		var comment ArticleComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = user.ID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case "momentComment":
		var comment MomentComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = user.ID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case "diaryComment":
		var comment DiaryComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = user.ID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	case "diaryBookComment":
		var comment DiaryBookComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = user.ID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			logrus.Info(err.Error())
		}
	}

	common.Response(c, "评论成功")
}

func commentBind(comment interface{}, c *fasthttp.RequestCtx) {
	err := common.BindWithJson(c, comment)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"controller": "moment",
		}).Info(err.Error())
		common.Response(c, "参数无效")
		return
	}
}
