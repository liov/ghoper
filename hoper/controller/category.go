package controller

import (
	"github.com/kataras/iris"
	"hoper/controller/common"
	"hoper/initialize"
	"hoper/model/e"
	"hoper/model/ov"
)

func GetCategory(c iris.Context) {
	var categories []ov.Category
	initialize.DB.Find(&categories)
	common.Response(c, categories, e.GetMsg(e.SUCCESS), e.SUCCESS)
}
