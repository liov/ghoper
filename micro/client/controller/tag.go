package controller

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/vo"
	"strconv"
)

func GetTags(c iris.Context) {

	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))

	var tags []vo.Tag

	err := initialize.DB.Select("name").
		Order("count desc").Limit(pageSize).Offset(pageNo).Find(&tags).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	common.Response(c, tags)
}

func GetTagTotal(maps interface{}) (count int) {
	initialize.DB.Model(&model.Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) *vo.Tag {
	if name == "" {
		return nil
	}
	var tag vo.Tag
	initialize.DB.Select("name,count").Where("name = ?", name).First(&tag)
	if tag.Name != "" {
		return &tag
	}

	return nil
}

func ExistMoodByName(name string) *vo.Mood {
	if name == "" {
		return nil
	}
	var mood vo.Mood
	initialize.DB.Select("name,count").Where("name = ?", name).First(&mood)
	if mood.Name != "" {
		return &mood
	}

	return nil
}

func AddTag(c iris.Context) bool {

	name := c.URLParam("name")
	userID := c.Values().Get("userID").(uint64)
	initialize.DB.Create(&model.Tag{
		Name:   name,
		UserID: userID,
	})

	return true
}

/*func ExistTagByID(id int) bool {
	var tag model.Tag
	initialize.DB.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}*/

func DeleteTag(id int) bool {
	initialize.DB.Where("id = ?", id).Delete(&model.Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	initialize.DB.Model(&model.Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func CleanAllTag() bool {
	initialize.DB.Unscoped().Where("deleted_on != ? ", 0).Delete(&model.Tag{})

	return true
}

const (
	flagTag = iota
	flagMood
	flagCategory
)

var IndexFlag = map[int8]string{
	flagTag:      "Tag",
	flagMood:     "Mood",
	flagCategory: "Category",
}

func setFlagCountToRedis(flag int8, name string, num int8) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	err := conn.Send("SELECT", 11)
	_, err = conn.Do("HINCRBY", IndexFlag[flag]+"_Count", name, num)
	if err != nil {
		golog.Error("缓存失败:", err)
	}
	return nil
}

func getFlagCountToRedis(flag int8) int64 {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	conn.Send("SELECT", 11)
	conn.Send("HGETALL", "Flag_Count")
	conn.Flush()
	conn.Receive()

	data, err := redis.Int64Map(conn.Receive())
	count := data[IndexFlag[flag]]

	if err != nil {
		golog.Error("缓存失败:", err)
	}
	return count
}
