package controller

import (
	"github.com/jinzhu/gorm"
	"hoper/utils/ulog"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"hoper/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/e"
	"hoper/model/ov"
)

type Comment interface {
	GetCommentType()
}

func AddComment(c iris.Context) {
	userID := c.Values().Get("userID").(uint64)
	kind := c.Params().Get("kind")
	refId, _ := c.Params().GetUint64("refId")
	if limitErr := common.Limit(model.CommentMinuteLimit,
		model.CommentMinuteLimitCount,
		model.CommentDayLimit,
		model.CommentMinuteLimitCount, userID); limitErr != nil {
		common.Response(c, limitErr.Error())
		return
	}

	nowTime := time.Now()
	var comment = func() interface{} {
		switch KindIndex[kind] {
		case kindArticle:
			return &model.ArticleComment{CreatedAt: nowTime, UserID: userID, ArticleID: refId, Status: 1}

		case kindMoment:
			return &model.MomentComment{CreatedAt: nowTime, UserID: userID, MomentID: refId, Status: 1}

		case kindDiary:
			return &model.DiaryComment{CreatedAt: nowTime, UserID: userID, DiaryID: refId, Status: 1}

		case kindDiaryBook:
			return &model.DiaryBookComment{CreatedAt: nowTime, UserID: userID, DiaryBookID: refId, Status: 1}
		}
		return nil
	}()

	commentBind(comment, c)
	if err := initialize.DB.Create(comment).Error; err != nil {
		ulog.Error(err)
	}

	setCountToRedis(userID, refId, KindIndex[kind], actionComment, 1)
	common.Response(c, "评论成功", e.SUCCESS)
}

func commentBind(comment interface{}, c iris.Context) {
	err := c.ReadJSON(comment)
	if err != nil {
		ulog.Error(err)
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
			ulog.Error(err)
		}
	case kindMoment:
		var comment ov.MomentComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		if err := initialize.DB.Create(&comment).Error; err != nil {
			ulog.Error(err)
		}
	case kindDiary:
		var comment ov.DiaryComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		if err := initialize.DB.Create(&comment).Error; err != nil {
			ulog.Error(err)
		}
	case kindDiaryBook:
		var comment ov.DiaryBookComment
		commentBind(&comment, ctx)
		comment.CreatedAt = nowTime
		if err := initialize.DB.Create(&comment).Error; err != nil {
			ulog.Error(err)
		}
	}
}

func GetComments(ctx iris.Context) {
	kind := ctx.Params().Get("kind")
	refId := ctx.Params().Get("refId")
	offset := ctx.URLParam("offset")
	limit, _ := strconv.Atoi(ctx.URLParam("limit"))
	rootID := ctx.URLParam("rootId")

	var db = func(c interface{}) interface{} {
		//var DB = initialize.DB.Where(kind+"_id = ? AND root_id = ?", refId, rootID).
		var DB = initialize.DB.Where(kind+"_id = ?", refId).
			Order("sequence DESC,approve_count DESC,created_at DESC").Preload("User")

		if rootID == "0" {
			DB = DB.Preload("SubComments", func(db *gorm.DB) *gorm.DB {
				return db.Limit(25).Offset(0).Order("sequence DESC,approve_count DESC,created_at DESC")
			}).Preload("SubComments.User")
		}

		if err := DB.Limit(limit).Offset(offset).Find(c).Error; err != nil {
			ulog.Error(err)
		}

		return c
	}

	var comments = func() interface{} {
		switch KindIndex[kind] {
		case kindArticle:
			c := make([]ov.ArticleComment, 0, limit)
			return db(&c)
		case kindMoment:
			c := make([]ov.MomentComment, 0, limit)
			return db(&c)
		case kindDiary:
			c := make([]ov.DiaryComment, 0, limit)
			return db(&c)
		case kindDiaryBook:
			c := make([]ov.DiaryBookComment, 0, limit)
			return db(&c)
		}
		return nil
	}()
	common.Res(ctx, iris.Map{
		"data": comments,
		//"comment_count": count,
		"code": 200,
	})
}
