package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	LoginError   = 1000 //用户名或密码错误
	LoginTimeout = 1001 //登录超时
	InActive     = 1002 //未激活账号
	WuQuanXian   = 1003

	ErrorExistTag        = 10001
	ErrorNotExistTag     = 10002
	ErrorNotExistArticle = 10003

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004

	// 保存图片失败
	ErrorUploadSaveImageFail = 30001
	// 检查图片失败
	ErrorUploadCheckImageFail = 30002
	// 校验图片错误，图片格式或大小有问题
	ErrorUploadCheckImageFormat = 30003
	//尝试次数过多
	TimeTooMuch = 40001
)

var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR:                       "fail",
	InvalidParams:               "请求参数错误",
	ErrorExistTag:               "已存在该标签名称",
	ErrorNotExistTag:            "该标签不存在",
	ErrorNotExistArticle:        "该文章不存在",
	ErrorAuthCheckTokenFail:     "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:  "Token已超时",
	ErrorAuthToken:              "Token生成失败",
	ErrorAuth:                   "Token错误",
	ErrorUploadSaveImageFail:    "保存图片失败",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadCheckImageFormat: "校验图片错误，图片格式或大小有问题",
	LoginError:                  "用户名或密码错误",
	LoginTimeout:                "登录超时",
	InActive:                    "未激活账号",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
