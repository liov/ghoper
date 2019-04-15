package controller

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/ov"
)

func GetTags(c iris.Context) {

	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))

	var tags []ov.Tag

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

func ExistTagByName(tag *ov.Tag, userID uint64) bool {
	if tag.Name == "" {
		return false
	}
	var nTag model.Tag
	initialize.DB.Select("name,count").Where("name = ?", tag.Name).First(&nTag)
	if nTag.Count > 0 {
		return true
	} else {
		newTag := model.Tag{Name: tag.Name, UserID: userID, CreatedAt: time.Now()}
		initialize.DB.Create(&newTag)
	}
	return true
}

func CreatCategory(category *ov.Category, userID uint64) uint64 {
	if category.Name == "" {
		return 0
	}
	newCategory := model.Category{Name: category.Name, Count: 1, UserID: userID, CreatedAt: time.Now()}
	initialize.DB.Create(&newCategory)
	return newCategory.ID
}

func CreatSerial(title *string, userID uint64) uint64 {
	if *title == "" {
		return 0
	}
	var newSerial model.Serial
	initialize.DB.Select("title,count").Where("title = ?", title).First(&newSerial)
	if newSerial.Title != "" {
		initialize.DB.Model(&newSerial).UpdateColumn("count", newSerial.Count+1)
	} else {
		newSerial = model.Serial{Title: *title, Count: 1, UserID: userID, CreatedAt: time.Now()}
		initialize.DB.Create(&newSerial)
	}

	return newSerial.ID
}

func ExistMoodByName(name string) *ov.Mood {
	if name == "" {
		return nil
	}
	newMood := model.Mood{Name: name, Count: 1, CreatedAt: time.Now()}
	initialize.DB.Create(&newMood)
	return &ov.Mood{Name: newMood.Name, Description: newMood.Description, ExpressionURL: newMood.ExpressionURL}
}

func AddTag(c iris.Context) bool {

	name := c.URLParam("name")
	userID := c.Values().Get("userID").(uint64)
	initialize.DB.Create(&model.Tag{
		Name:      name,
		UserID:    userID,
		CreatedAt: time.Now(),
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
