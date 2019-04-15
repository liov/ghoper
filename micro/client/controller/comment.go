package controller

import (
	"strconv"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
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
	refId, _ := c.Params().GetUint64("ref_id")
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
		golog.Error(err)
	}

	setCountToRedis(userID, refId, KindIndex[kind], actionComment, 1)
	common.Response(c, "评论成功", e.SUCCESS)
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
	refId := ctx.Params().Get("ref_id")
	pageNo, _ := strconv.Atoi(ctx.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(ctx.URLParam("pageSize"))
	rootID := ctx.URLParam("root_id")
	var count int64
	/*
		var comments = func() interface{} {
			switch KindIndex[kind] {
			case kindArticle:
				return make([]ov.ArticleComment, 0, pageSize)

			case kindMoment:
				return make([]ov.MomentComment, 0, pageSize)

			case kindDiary:
				return make([]ov.DiaryComment, 0, pageSize)

			case kindDiaryBook:
				return make([]ov.DiaryBookComment,0, pageSize)
			}
			return nil
		}()*/
	var where string
	if rootID == "0" {
		where = "root_id = id"
	} else {
		where = "root_id = '" + rootID + "'"
	}

	switch KindIndex[kind] {
	case kindArticle:
		var c []ov.ArticleComment
		if err := initialize.DB.Where(kind+"_id = ? AND "+where, refId).Order("sequence desc,created_at desc").Limit(pageSize).
			Offset(pageNo * pageSize).Preload("User").Find(&c).Count(&count).Error; err != nil {
			golog.Error(err)
		}
		common.Response(ctx, c)
		return
	case kindMoment:
		var c []ov.MomentComment
		if err := initialize.DB.Where(kind+"_id = ? AND "+where, refId).Order("sequence desc,created_at desc").Limit(pageSize).
			Offset(pageNo * pageSize).Preload("User").Find(&c).Count(&count).Error; err != nil {
			golog.Error(err)
		}
		common.Response(ctx, c)
		return
	case kindDiary:
		var c []ov.DiaryComment
		if err := initialize.DB.Where(kind+"_id = ? AND "+where, refId).Order("sequence desc,created_at desc").Limit(pageSize).
			Offset(pageNo * pageSize).Preload("User").Find(&c).Count(&count).Error; err != nil {
			golog.Error(err)
		}
		common.Response(ctx, c)
		return
	case kindDiaryBook:
		var c []ov.DiaryBookComment
		if err := initialize.DB.Where(kind+"_id = ? AND "+where, refId).Order("sequence desc,created_at desc").Limit(pageSize).
			Offset(pageNo * pageSize).Preload("User").Find(&c).Count(&count).Error; err != nil {
			golog.Error(err)
		}
		common.Response(ctx, c)
		return
	}
}
