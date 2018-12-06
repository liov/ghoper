package controller

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"service/controller/common"
	"service/model"
	"service/utils"
	"time"
)

type ArticleComment struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	UserID      uint      `json:"user_id"`
	Content     string    `gorm:"type:varchar(500)" json:"content"`
	HTMLContent string    `gorm:"type:varchar(500)" json:"html_content"`
	ContentType int       `json:"content_type"`
	ArticleID   uint      `json:"article_id"` //话题或投票的ID
	ParentID    uint      `json:"parent_id"`  //直接父评论的ID
}

func (a *ArticleComment) GetCommentType() {

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

func (m *MomentComment) GetCommentType() {

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

func (d *DiaryComment) GetCommentType() {

}

type DiaryBookComment struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	UserID      uint      `json:"user_id"`
	Content     string    `gorm:"type:varchar(500)" json:"content"`
	DiaryBookID uint      `json:"diary_book_id"` //瞬间ID
	ParentID    uint      `json:"parent_id"`     //直接父评论的ID
}

func (d *DiaryBookComment) GetCommentType() {

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
	classify := utils.ToSting(c.QueryArgs().Peek("classify"))
	var err error
	switch classify {
	case "articleComment":
		var comment ArticleComment
		err = common.BindWithJson(c, &comment)
	case "momentComment":
		var comment MomentComment
		err = common.BindWithJson(c, &comment)
	case "diaryComment":
		var comment DiaryComment
		err = common.BindWithJson(c, &comment)
	case "diaryBookComment":
		var comment DiaryBookComment
		err = common.BindWithJson(c, &comment)
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"controller": "moment",
		}).Info(err.Error())
		common.Response(c, "参数无效")
		return
	}
	common.Response(c, "评论成功")
}
