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
	user := c.GetViewData()["user"].(User)
	if limitErr := common.Limit(model.CommentMinuteLimit,
		model.CommentMinuteLimitCount,
		model.CommentDayLimit,
		model.CommentMinuteLimitCount, user.ID); limitErr != nil {
		common.Response(c, limitErr.Error())
		return
	}
	classify := c.GetViewData()["classify"].(string)
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
