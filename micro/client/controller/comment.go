package controller

import (
	"github.com/jinzhu/gorm"
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

	var where string
	if rootID == "0" {
		where = "root_id = id"
	} else {
		where = "root_id = '" + rootID + "'"
	}

	//var comments = make([]ov.MomentComment, 0, pageSize) 是可以的
	var db = func(c interface{}) interface{} {
		if err := initialize.DB.Where(kind+"_id = ? AND "+where, refId).Order("sequence desc,created_at desc").Limit(pageSize).
			Offset(pageNo*pageSize).Preload("User").Preload("SubComments", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(5).Offset(0)
		}).Preload("SubComments.User").Find(c).Count(&count).Error; err != nil {
			golog.Error(err)
		}
		return c
	}

	var comments = func() interface{} {
		switch KindIndex[kind] {
		case kindArticle:
			c := make([]ov.ArticleComment, 0, pageSize)
			return db(&c)
		case kindMoment:
			c := make([]ov.MomentComment, 0, pageSize)
			return db(&c)
		case kindDiary:
			c := make([]ov.DiaryComment, 0, pageSize)
			return db(&c)
		case kindDiaryBook:
			c := make([]ov.DiaryBookComment, 0, pageSize)
			return db(&c)
		}
		return nil
	}()
	common.Res(ctx, iris.Map{
		"data":          comments,
		"comment_count": count,
		"code":          200,
	})
}
