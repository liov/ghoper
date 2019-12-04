package common

import (
	"github.com/kataras/iris"
	"hoper/model/e"
	"hoper/utils/ulog"
)

type H map[string]interface{}

type ResData struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
}

//先信息后数据最后状态码
//入参1. Data interface{},Msg string,Code int
//2.Msg,Code |data默认nil
//3.Data,Msg |code默认SUCCESS
//4.Msg |data默认nil code默认ERROR
//5.Data |msg默认"",code默认SUCCESS
func Response(ctx iris.Context, res ...interface{}) {

	var resData ResData

	if len(res) == 1 {
		resData.Code = e.ERROR
		if msgTmp, ok := res[0].(string); ok {
			resData.Msg = msgTmp
			resData.Data = nil
		} else {
			resData.Data = res[0]
			resData.Code = e.SUCCESS
		}
	} else if len(res) == 2 {
		if msgTmp, ok := res[0].(string); ok {
			resData.Data = nil
			resData.Msg = msgTmp
			resData.Code = res[1].(int)
		} else {
			resData.Data = res[0]
			resData.Msg = res[1].(string)
			resData.Code = e.SUCCESS
		}
	} else {
		resData.Data = res[0]
		resData.Msg = res[1].(string)
		resData.Code = res[2].(int)
	}

	num, err := ctx.JSON(resData)

	if err != nil {
		ulog.Error(num, err)
	}
}

/*func Response(ctx iris.Context,Data interface{},Msg string,Code int){
	num, err := ctx.JSON(iris.Map{
		"Code": Code,
		"Msg":  Msg,
		"Data": Data,
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
