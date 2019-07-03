package common

import (
	"github.com/kataras/iris"
	"hoper/model/e"
	"hoper/utils/ulog"
)

type H map[string]interface{}

type ResData struct {
	data interface{}
	msg string
	code int
}
//先信息后数据最后状态码
//入参1. data interface{},msg string,code int
//2.msg,code |data默认nil
//3.data,msg |code默认SUCCESS
//4.msg |data默认nil code默认ERROR
//5.data |msg默认"",code默认SUCCESS
func Response(ctx iris.Context, res ...interface{}) {

	var resData ResData

	if len(res) == 1 {
		resData.code = e.ERROR
		if msgTmp, ok := res[0].(string); ok {
			resData.msg = msgTmp
			resData.data = nil
		} else {
			resData.data = res[0]
			resData.code = e.SUCCESS
		}
	} else if len(res) == 2 {
		if msgTmp, ok := res[0].(string); ok {
			resData.data = nil
			resData.msg = msgTmp
			resData.code = res[1].(int)
		} else {
			resData.data = res[0]
			resData.msg = res[1].(string)
			resData.code = e.SUCCESS
		}
	} else {
		resData.data = res[0]
		resData.msg = res[1].(string)
		resData.code = res[2].(int)
	}

	num, err := ctx.JSON(&resData)

	if err != nil {
		ulog.Error(num, err)
	}
}

/*func Response(ctx iris.Context,data interface{},msg string,code int){
	num, err := ctx.JSON(iris.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	if err != nil {
		ulog.Error(num, err)
	}
}*/

func Res(c iris.Context, h iris.Map) {
	num, err := c.JSON(h)
	if err != nil {
		ulog.Error(num, err)
	}
}
