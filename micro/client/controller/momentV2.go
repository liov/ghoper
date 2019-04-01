package controller

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"hoper/client/controller/cachekey"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/e"
	"hoper/model/vo"
	"hoper/utils"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/3/29
 * @description：
 */
//DTO

func AddMoment(c iris.Context) {

	user := c.Values().Get("user").(vo.User)

	//Limit这个函数的封装呢，费了点功夫，之前的返回值想到用err，不过在sendErr这出了点问题，决定返回值改用string，这样是不规范的
	if limitErr := common.Limit(model.MomentMinuteLimit,
		model.MomentMinuteLimitCount,
		model.MomentDayLimit,
		model.MomentMinuteLimitCount, user.ID); limitErr != nil {
		common.Response(c, limitErr.Error(), e.TimeTooMuch)
		return
	}

	var moment vo.Moment

	if err := c.ReadJSON(&moment); err != nil {
		logrus.WithFields(logrus.Fields{
			"model": "moment",
		}).Info(err.Error())
		common.Response(c, "参数无效")
		return
	}

	if utf8.RuneCountInString(moment.Content) > 500 {
		common.Response(c, "文章内容不能小于20个字")
		return
	}
	/*	moodName := moment.MoodName

		var mood model.Mood

		moodErr :=initialize.DB.Where("name = ?", moodName).Find(&mood).Error

		if moodErr != nil{
			mood.Name = moodName
			initialize.DB.Create(&mood)
			moment.Mood = mood
		} else {
			moment.Mood = mood
		}*/

	nowTime := time.Now()
	moment.CreatedAt = nowTime
	//moment.Mood = Mood{Name: moment.MoodName}

	if err := validationMoment(c, &moment); err != nil {
		return
	}
	moment.UserID = user.ID
	moment.BrowseCount = 1
	/*	moment.Status = model.ArticleVerifying
		moment.ModifyTimes = 0
		moment.ParentID = 0 */
	/*	user.Score = user.Score + model.ArticleScore
		user.ArticleCount = user.ArticleCount + 1

		if model.UserToRedis(user) != nil {
			common.SendErr(c,"error")
			return
		}*/
	moment.Content = strings.TrimSpace(moment.Content)
	var mood *vo.Mood
	if mood = ExistMoodByName(moment.MoodName); mood != nil {
		setFlagCountToRedis(flagTag, moment.MoodName, 1)
	} else if mood.Name != "" {
		newMood := model.Mood{CreatedAt: nowTime, Name: moment.MoodName, Count: 1}
		initialize.DB.Create(&newMood)
		moment.Mood = vo.Mood{Name: newMood.Name, Description: newMood.Description, ExpressionURL: newMood.ExpressionURL}
	}

	saveErr := initialize.DB.Create(&moment).Error

	for _, v := range moment.Tags {
		var tag *vo.Tag
		if tag = ExistTagByName(v.Name); tag != nil {
			setFlagCountToRedis(flagTag, v.Name, 1)
		} else if tag.Name != "" {
			newTag := model.Tag{CreatedAt: nowTime, Name: v.Name, Count: 1}
			initialize.DB.Create(&newTag)
		}
		momentTag := model.MomentTag{MomentID: moment.ID, TagName: v.Name}
		initialize.DB.Create(&momentTag)
	}

	if saveErr != nil {
		common.Response(c, "创建出错")
		return
	}

	//var moments []model.Moment
	moment.User = user
	moment.Mood = *mood
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	/*	if moment.DescFlag == 0 {
			value, _ := json.Marshal(moment)
			_, err := conn.Do("LPUSH", gredis.Moments, value)
			if err != nil {
				return
			}
		} else {
			value, _ := json.Marshal(moment)
			_, err := conn.Do("LPUSH", gredis.TopMoments, value)
			if err != nil {
				return
			}
		}*/

	value, _ := utils.Json.MarshalToString(moment)
	conn.Send("SELECT", kindMoment)
	conn.Send("LPUSH", cachekey.Moments+"_V2", value)
	_, err := conn.Do("INCR", cachekey.Moments+"_Count")
	if err != nil {
		return
	}

	common.Response(c, "新建成功", e.SUCCESS)
}

func GetMomentsV2(c iris.Context) {
	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))
	//l := list.New()
	userID := c.Values().Get("userID").(uint64)
	key := cachekey.Moments + "_V2"

	var moments []vo.Moment
	var userAction *UserAction
	if userID > 0 {
		userAction = getRedisAction(strconv.FormatUint(userID, 10), kindMoment)
	}
	var count, topCount int64
	if moments, count, topCount = getRedisMomentsV2(key, pageNo, pageSize); moments != nil {

		common.Res(c, iris.Map{"data": moments,
			"count":       count,
			"top_count":   topCount,
			"user_action": userAction,
			"msg":         e.GetMsg(e.SUCCESS),
			"code":        e.SUCCESS})
		return
	}

	err := initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Preload("User").Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count").
		Order("sequence desc,id desc").Limit(pageSize).
		Offset(pageNo * pageSize).Find(&moments).Count(&count).Error
	err = initialize.DB.Model(vo.Moment{}).Where("sequence = ?", 9).Count(&topCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if userID > 0 {
		getRedisAction(strconv.FormatUint(uint64(userID), 10), kindMoment)
	}

	common.Res(c, iris.Map{"data": moments,
		"count":       count,
		"top_count":   topCount,
		"user_action": userAction,
		"msg":         e.GetMsg(e.SUCCESS),
		"code":        e.SUCCESS})

	setRedisMomentsV2(key, moments, count, topCount)

}

func validationMoment(c iris.Context, moment *vo.Moment) (err error) {

	err = &e.ValidtionError{Msg: "参数无效"}

	if moment.Content == "" || utf8.RuneCountInString(moment.Content) <= 0 {
		common.Response(c, "文章内容不能为空")
		return
	}

	if utf8.RuneCountInString(moment.Content) > model.MaxContentLen {
		msg := "文章内容不能超过" + strconv.Itoa(model.MaxContentLen) + "个字符"
		common.Response(c, msg)
		return
	}

	/*	if moment.Tags == nil || len(moment.Tags) <= 0 {
		SendErrJSON(c,"请选择标签")
		return
	}*/

	return nil
}

func getRedisMomentsV2(key string, pageNo int, PageSize int) ([]vo.Moment, int64, int64) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	var moments []vo.Moment
	conn.Send("SELECT", kindMoment)
	if exist, err := redis.Bool(conn.Do("EXISTS", key)); !exist || err != nil {
		return nil, 0, 0
	}
	start := pageNo * PageSize

	data, _ := redis.Strings(conn.Do("LRANGE", key, start, start+PageSize-1))
	for _, mv := range data {
		var moment vo.Moment
		utils.Json.UnmarshalFromString(mv, &moment)
		conn.Send("HINCRBY", strings.Join([]string{IndexKind[kindMoment], strconv.FormatUint(moment.ID, 10), "Action", "Count"}, "_"), IndexAction[actionBrowse], 1)
		actionCount := getActionCount(moment.ID, kindMoment)
		actionCount.BrowseCount = actionCount.BrowseCount + 1
		moment.ActionCount = *actionCount
		moments = append(moments, moment)
	}
	conn.Do("")
	conn.Send("GET", "Moment_List_Count")
	conn.Send("GET", "Moment_List_Top_Count")
	conn.Flush()
	count, _ := redis.Int64(conn.Receive())
	topCount, _ := redis.Int64(conn.Receive())
	return moments, count, topCount
}

func setRedisMomentsV2(key string, moments []vo.Moment, count int64, topCount int64) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("SELECT", kindMoment)
	for _, mv := range moments {
		mv.BrowseCount = mv.BrowseCount + 1
		//mv.Index = mi
		value, _ := utils.Json.MarshalToString(mv)
		err := conn.Send("RPUSH", key, value)
		if err != nil {
			return err
		}
	}
	conn.Send("SET", "Moment_List_Count", strconv.FormatInt(count, 10))
	conn.Do("SET", "Moment_List_Top_Count", strconv.FormatInt(topCount, 10))
	return nil
}

func GetMomentV2(c iris.Context) {

	index := c.URLParam("index")
	userID := c.Values().Get("userID").(uint64)
	var userAction *UserAction
	if userID > 0 {
		userAction = getRedisAction(strconv.FormatUint(userID, 10), kindMoment)
	}
	if moment := getRedisMomentV2(index); moment != nil {

		common.Res(c, iris.Map{"data": moment,
			"user_action": userAction,
			"msg":         e.GetMsg(e.SUCCESS),
			"code":        e.SUCCESS})
		return
	}

	id, err := c.Params().GetUint64("id")
	var moment vo.Moment
	err = initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count,permission").
		Where("id = ?", id).First(&moment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		golog.Error(err)
		return
	}
	common.Res(c, iris.Map{"data": moment,
		"user_action": userAction,
		"msg":         e.GetMsg(e.SUCCESS),
		"code":        e.SUCCESS})
}

func getRedisMomentV2(index string) *vo.Moment {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	key := cachekey.Moments + "_V2"
	conn.Send("SELECT", kindMoment)

	data, err := redis.String(conn.Do("LINDEX", key, index))
	var moment vo.Moment
	err = utils.Json.UnmarshalFromString(data, &moment)
	conn.Do("HINCRBY", strings.Join([]string{IndexKind[kindMoment], strconv.FormatUint(moment.ID, 10), "Action", "Count"}, "_"), IndexAction[actionBrowse], 1)
	actionCount := getActionCount(moment.ID, kindMoment)
	moment.ActionCount = *actionCount
	if err != nil {
		golog.Error(err)
		return nil
	}
	return &moment

}
