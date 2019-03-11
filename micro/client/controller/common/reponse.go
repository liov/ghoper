package common

import (
	"fmt"
	"github.com/kataras/iris"
	"hoper/client/controller/common/e"
)

type H map[string]interface{}

//先信息后数据最后状态码
//入参 data interface{},msg string,code int
//data,msg code默认SUCCESS
//msg data默认nil code默认ERROR
//data msg默认"",code默认SUCCESS
//msg,code data默认nil
func Response(ctx iris.Context, res ...interface{}) {

	var msg string
	var code int
	var data interface{}
	if len(res) == 1 {
		code = e.ERROR
		if msgTmp, ok := res[0].(string); ok {
			msg = msgTmp
			data = nil
		} else {
			data = res[0]
			code = e.SUCCESS
		}
	} else if len(res) == 2 {
		if msgTmp, ok := res[0].(string); ok {
			data = nil
			msg = msgTmp
			code = res[1].(int)
		} else {
			data = res[0]
			msg = res[1].(string)
			code = e.SUCCESS
		}
	} else {
		data = res[0]
		msg = res[1].(string)
		code = res[2].(int)
	}

	num, err := ctx.JSON(iris.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	if err != nil {
		fmt.Println(num, err)
	}
}

func Res(c iris.Context, h iris.Map) {
	num, err := c.JSON(h)
	if err != nil {
		fmt.Println(num, err)
	}
}
