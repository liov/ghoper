package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/e"
	"hoper/model/ov"
	"strconv"

	"time"
)

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
		var comment model.ArticleComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	case kindMoment:
		var comment model.MomentComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	case kindDiary:
		var comment model.DiaryComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	case kindDiaryBook:
		var comment model.DiaryBookComment
		commentBind(&comment, c)
		comment.CreatedAt = nowTime
		comment.UserID = userID
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	}

	common.Response(c, "评论成功")
}

func commentBind(comment interface{}, c iris.Context) {
	err := c.ReadJSON(comment)
	if err != nil {
		golog.Error(err)
		common.Response(c, "参数无效")
		return
	}
}

func GetComment(ctx iris.Context) {
	kind := ctx.Params().Get("kind")
	nowTime := time.Now()
	switch KindIndex[kind] {
	case kindArticle:
		var comment ov.ArticleComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	case kindMoment:
		var comment ov.MomentComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	case kindDiary:
		var comment ov.DiaryComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	case kindDiaryBook:
		var comment ov.DiaryBookComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		if err := initialize.DB.Create(&comment).Error; err != nil {
			golog.Error(err)
		}
	}
}

func GetComments(ctx iris.Context) {
	kind := ctx.Params().Get("kind")
	pageNo, _ := strconv.Atoi(ctx.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(ctx.URLParam("pageSize"))
	var count int64

	var comments = func() interface{} {
		switch KindIndex[kind] {
		case kindArticle:
			return make([]ov.ArticleComment, 0, pageSize)

		case kindMoment:
			return make([]ov.MomentComment, 0, pageSize)

		case kindDiary:
			return make([]ov.DiaryComment, 0, pageSize)

		case kindDiaryBook:
			return make([]ov.DiaryBookComment, 0, pageSize)
		}
		return nil
	}()

	if err := initialize.DB.Order("sequence desc,id desc").Limit(pageSize).
		Offset(pageNo * pageSize).Find(&comments).Count(&count).Error; err != nil {
		golog.Error(err)
	}
	common.Response(ctx, comments, e.SUCCESS)
}
