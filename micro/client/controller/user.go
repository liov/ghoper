package controller

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/client/controller/credis"
	"hoper/client/controller/upload"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/e"
	"hoper/model/ov"
	"hoper/utils"
	"hoper/utils/mail"
)

const (
	activeDuration = 24 * 60 * 60
	resetDuration  = 24 * 60 * 60
)

type User struct {
	ov.User
	ov.KindOwnCount
	Role uint8 `gorm:"type:smallint;default:0" json:"-"` //管理员or用户
}

func sendMail(action string, title string, curTime int64, user model.User) {
	siteName := initialize.Config.Server.SiteName
	siteURL := "https://" + initialize.Config.Server.Host
	secretStr := strconv.FormatInt(curTime, 10) + user.Email + user.Password
	secretStr = fmt.Sprintf("%x", md5.Sum(utils.ToBytes(secretStr)))
	actionURL := siteURL + "/user" + action + "/"

	actionURL = actionURL + strconv.FormatUint(user.ID, 10) + "/" + secretStr
	golog.Info(actionURL)
	content := "<p><b>亲爱的" + user.Name + ":</b></p>" +
		"<p>我们收到您在 " + siteName + " 的注册信息, 请点击下面的链接, 或粘贴到浏览器地址栏来激活帐号.</p>" +
		"<a href=\"" + actionURL + "\">" + actionURL + "</a>" +
		"<p>如果您没有在 " + siteName + " 填写过注册信息, 说明有人滥用了您的邮箱, 请删除此邮件, 我们对给您造成的打扰感到抱歉.</p>" +
		"<p>" + siteName + " 谨上.</p>"

	if action == "/reset" {
		content = "<p><b>亲爱的" + user.Name + ":</b></p>" +
			"<p>你的密码重设要求已经得到验证。请点击以下链接, 或粘贴到浏览器地址栏来设置新的密码: </p>" +
			"<a href=\"" + actionURL + "\">" + actionURL + "</a>" +
			"<p>感谢你对" + siteName + "的支持，希望你在" + siteName + "的体验有益且愉快。</p>" +
			"<p>(这是一封自动产生的email，请勿回复。)</p>"
	}
	//content += "<p><img src=\"" + siteURL + "/images/logo.png\" style=\"height: 42px;\"/></p>"
	//fmt.Println(content)

	mail.SendMail(user.Email, title, content)
}
func verifyLink(cacheKey string, c iris.Context) (model.User, error) {
	var user model.User

	userID, _ := c.Params().GetInt("id")
	if userID <= 0 {
		return user, errors.New("无效的链接")
	}
	secret := c.Params().Get("secret")
	if secret == "" {
		return user, errors.New("无效的链接")
	}
	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	emailTime, redisErr := redis.Int64(RedisConn.Do("GET", cacheKey+strconv.Itoa(userID)))
	if redisErr != nil {
		return user, errors.New("无效的链接")
	}

	if err := initialize.DB.First(&user, userID).Error; err != nil {
		return user, errors.New("无效的链接")
	}

	secretStr := strconv.Itoa((int)(emailTime)) + user.Email + user.Password

	secretStr = fmt.Sprintf("%x", md5.Sum(utils.ToBytes(secretStr)))

	if secret != secretStr {
		return user, errors.New("无效的链接")
	}
	return user, nil
}

// ActiveSendMail 发送激活账号的邮件
func ActiveSendMail(c iris.Context) {

	// 接收到的email参数是加密后的，不能加email验证规则
	type ReqData struct {
		Email string `json:"email" binding:"required"`
	}

	var reqData ReqData
	// 只接收一个email参数
	if err := c.ReadJSON(&reqData); err != nil {
		common.Response(c, "参数无效", e.InvalidParams)
		return
	}

	var user model.User
	user.Email = reqData.Email

	var decodeBytes []byte
	var decodedErr error
	if decodeBytes, decodedErr = base64.StdEncoding.DecodeString(user.Email); decodedErr != nil {
		common.Response(c, "参数无效", e.InvalidParams)
		return
	}
	user.Email = *(*string)(unsafe.Pointer(&decodeBytes))

	if err := initialize.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		common.Response(c, "无效的邮箱")
		return
	}

	curTime := time.Now().Unix()
	activeUser := model.ActiveTime + strconv.FormatUint(user.ID, 10)

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("SET", activeUser, curTime, "EX", activeDuration); err != nil {
		golog.Error("redis set failed:", err)
	}
	go func() {
		sendMail("/active", "账号激活", curTime, user)
	}()

	common.Res(c, iris.Map{"email": user.Email})

}

// ActiveAccount 激活账号
func ActiveAccount(c iris.Context) {
	var err error
	var user model.User
	if user, err = verifyLink(model.ActiveTime, c); err != nil {
		common.Response(c, "激活链接已失效")
		return
	}

	updatedData := map[string]interface{}{
		"status":       model.UserStatusActived,
		"activated_at": time.Now(),
	}

	if err := initialize.DB.Model(&user).Updates(updatedData).Error; err != nil {
		common.Response(c, "创建出错", e.ERROR)
		return
	}

	initialize.DB.Create(&model.Favorites{Name: "默认收藏夹", UserID: user.ID, Status: 1, Count: 0, CreatedAt: time.Now()})

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("DEL", model.ActiveTime+strconv.FormatUint(user.ID, 10)); err != nil {
		golog.Info(err)
	}
	common.Response(c, user.Email, "激活成功", e.SUCCESS)
}

// ResetPasswordMail 发送重置密码的邮件
func ResetPasswordMail(c iris.Context) {

	type UserReqData struct {
		Email       string `json:"email" binding:"required,email"`
		LuosimaoRes string `json:"luosimaoRes"`
	}
	var userData UserReqData
	if err := c.ReadJSON(&userData); err != nil {
		common.Response(c, "无效的邮箱")
		return
	}

	verifyErr := utils.LuosimaoVerify(initialize.Config.Server.LuosimaoVerifyURL, initialize.Config.Server.LuosimaoAPIKey, userData.LuosimaoRes)

	if verifyErr != nil {
		common.Response(c, verifyErr.Error())
		return
	}

	var user model.User
	if err := initialize.DB.Where("email = ?", userData.Email).Find(&user).Error; err != nil {
		common.Response(c, "没有邮箱为 "+userData.Email+" 的用户")
		return
	}

	curTime := time.Now().Unix()
	resetUser := model.ResetTime + strconv.FormatUint(user.ID, 10)

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("SET", resetUser, curTime, "EX", resetDuration); err != nil {
		golog.Error("redis set failed:", err)
	}
	go func() {
		sendMail("/ac", "修改密码", curTime, user)
	}()

	common.Res(c, iris.Map{"data": "修改成功"})
}

// VerifyResetPasswordLink 验证重置密码的链接是否失效
func VerifyResetPasswordLink(c iris.Context) {

	if _, err := verifyLink(model.ResetTime, c); err != nil {
		golog.Error(err)
		common.Response(c, "重置链接已失效")
		return
	}
	common.Response(c, "链接已发送")
}

// ResetPassword 重置密码
func ResetPassword(c iris.Context) {

	type UserReqData struct {
		Password string `json:"password" binding:"required,min=6,max=20"`
	}
	var userData UserReqData

	if err := c.ReadJSON(&userData); err != nil {
		common.Response(c, "参数无效")
		return
	}

	var verifErr error
	var user model.User
	if user, verifErr = verifyLink(model.ResetTime, c); verifErr != nil {
		common.Response(c, "重置链接已失效")
		return
	}

	user.Password = encryptPassword(userData.Password, userData.Password[0:5])

	if user.ID <= 0 {
		common.Response(c, "重置链接已失效")
		return
	}
	if err := initialize.DB.Model(&user).Update("pass", user.Password).Error; err != nil {
		common.Response(c, "error")
		return
	}

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("DEL", model.ResetTime+strconv.FormatUint(user.ID, 10)); err != nil {
		golog.Error("redis delelte failed:", err)
	}
}

// Login 用户登录
func Login(c iris.Context) {

	type Login struct {
		//Email string `json:"email" binding:"email"`
		//Phone	string `json:"phone"`
		Input    string `json:"input" binding:"required"`
		Password string `json:"password" binding:"required,min=6,max=20"`
		Luosimao string `json:"luosimao"`
	}

	var login Login

	var loginInput, password, luosimao, sql string

	if err := c.ReadJSON(&login); err != nil {
		common.Response(c, "账号或密码错误")
		return
	}
	if login.Input != "" {
		emailMatch, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+.[a-zA-Z]{2,3}$`, login.Input)

		phoneMatch, _ := regexp.MatchString(`^1[0-9]{10}$`, login.Input)
		if emailMatch {
			sql = "email = ?"
		} else if phoneMatch {
			sql = "phone = ?"
		} else {
			common.Response(c, "账号错误")
			return
		}
	}

	loginInput = login.Input
	password = login.Password
	luosimao = login.Luosimao

	verifyErr := utils.LuosimaoVerify(initialize.Config.Server.LuosimaoVerifyURL, initialize.Config.Server.LuosimaoAPIKey, luosimao)

	if verifyErr != nil {
		common.Response(c, verifyErr.Error())
		return
	}

	var user model.User
	if err := initialize.DB.Where(sql, loginInput).Find(&user).Error; err != nil {
		common.Response(c, "账号不存在", e.ERROR)
		return
	}

	if checkPassword(password, &user) {
		if user.Status == model.UserStatusInActive {
			//没看懂
			encodedEmail := base64.StdEncoding.EncodeToString(utils.ToBytes(user.Email))
			common.Res(c, iris.Map{"email": encodedEmail, "msg": "账号未激活,请进去邮箱点击激活"})
			return
		}

		tokenString, err := GenerateToken(&user)
		if err != nil {
			common.Response(c, "内部错误")
			return
		}

		initialize.DB.Model(&user).UpdateColumn("last_activated_at", time.Now())

		if err := UserToRedis(&User{
			User: ov.User{
				ID:        user.ID,
				Name:      user.Name,
				Sex:       user.Sex,
				Score:     user.Score,
				Signature: user.Signature,
				AvatarURL: user.AvatarURL,
				Status:    user.Status,
			},
			Role: user.Role,
		}); err != nil {
			common.Response(c, "内部错误.")
			return
		}

		c.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Path:     "/",
			Domain:   "hoper.xyz",
			Expires:  time.Now().Add(time.Duration(initialize.Config.Server.TokenMaxAge) * time.Second),
			MaxAge:   int(time.Duration(initialize.Config.Server.TokenMaxAge) * time.Second),
			Secure:   false,
			HttpOnly: true,
		})

		c.ResponseWriter().Header().Add("Authorization",tokenString)

		/*		session := sessions.Default(c)
				session.Set("user", user)
				session.Save()*/
		//userBytes, err := json.Marshal(user)
		//c.SetCookie("user", string(userBytes), initialize.ServerSettings.TokenMaxAge, "/", "hoper.xyz", false, true)
		common.Res(c, iris.Map{
			"token": tokenString,
			"data":  user,
			"msg":   "登录成功",
			"code":  e.SUCCESS,
		})

		return
	}
	common.Response(c, "账号或密码错误")
}

// Signup 用户注册
func Signup(c iris.Context) {

	type RegisterUser struct {
		Name     string  `json:"name" binding:"required,min=3,max=20"`
		Email    string  `json:"email" binding:"required,email"`
		Password string  `json:"password" binding:"required,min=6,max=20"`
		Phone    *string `json:"phone"`
		Sex      string  `json:"sex"`
	}

	var registerUser RegisterUser

	if err := c.ReadJSON(&registerUser); err != nil {
		common.Response(c, "参数无效")
		return
	}
	registerUser.Name = utils.AvoidXSS(registerUser.Name)
	registerUser.Name = strings.TrimSpace(registerUser.Name)
	registerUser.Email = strings.TrimSpace(registerUser.Email)

	if strings.Index(registerUser.Name, "@") != -1 {
		common.Response(c, "用户名中不能含有@字符")
		return
	}

	var user model.User
	if err := initialize.DB.Where("email = ? OR phone = ?", registerUser.Email, registerUser.Phone).Find(&user).Error; err == nil {
		if user.Phone != nil && *user.Phone == *registerUser.Phone {
			common.Response(c, "手机号已被注册")
			return
		} else if user.Email == registerUser.Email {
			common.Response(c, "邮箱已存在")
			return
		}
	}

	var newUser model.User
	nowTime := time.Now()
	newUser.CreatedAt = nowTime
	newUser.Name = registerUser.Name
	newUser.Sex = registerUser.Sex
	newUser.Email = registerUser.Email
	newUser.Phone = registerUser.Phone
	newUser.Password = encryptPassword(registerUser.Password, registerUser.Password)
	newUser.AvatarURL = "https://hoper.xyz/static/images/6cbeb5c8-7160-4b6f-a342-d96d3c00367a.jpg"
	//newUser.Role = model.UserRoleNormal
	newUser.Status = model.UserStatusInActive

	if err := initialize.DB.Create(&newUser).Error; err != nil {
		common.Response(c, err.Error(), e.ERROR)
		return
	}

	curTime := nowTime.Unix()
	activeUser := model.ActiveTime + strconv.FormatUint(newUser.ID, 10)

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("SET", activeUser, curTime, "EX", activeDuration); err != nil {
	}

	go func() {
		sendMail("/active", "账号激活", curTime, newUser)
	}()

	common.Response(c, newUser, "注册成功", e.SUCCESS)
}

// Logout 退出登录
func Logout(c iris.Context) {
	userID := c.Values().Get("userID").(uint64)
	initialize.DB.Model(&User{User: ov.User{ID: userID}}).UpdateColumn("last_activated_at", time.Now())

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("DEL", model.LoginUser+strconv.FormatUint(userID, 10)); err != nil {
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    "del",
		Path:     "/",
		Domain:   "hoper.xyz",
		Expires:  time.Now().Add(-1),
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	})

	common.Response(c, "已注销", e.SUCCESS)
}

func GetUserSelf(c iris.Context) {
	userID := c.Values().Get("userID").(uint64)
	var user model.User
	user.ID = userID
	initialize.DB.Preload("EduExps").Preload("WorkExps").Find(&user)
	//*string改值的方法
	phone := (*(user.Phone))[0:3] + "XXXX" + (*(user.Phone))[7:]
	user.Phone = &phone
	common.Response(c, user, e.GetMsg(e.SUCCESS), e.SUCCESS)
}

func LoginFlag(c iris.Context) {
	user := c.Values().Get("user").(*User)
	common.Response(c, user)
}

func GetUser(c iris.Context) {
	id := c.Params().GetUint64Default("id", 0)
	var user ov.User
	initialize.DB.Where("id=?", id).First(&user)
	common.Response(c, user, e.GetMsg(e.SUCCESS), e.SUCCESS)
}

func UpdateUser(c iris.Context) {
	var user, nUser model.User
	if err := c.ReadJSON(&nUser); err != nil {
		common.Response(c, "参数无效")
		return
	}
	userID := c.Values().Get("userID").(uint64)

	tx := initialize.DB.Begin()

	for _, v := range nUser.EduExps {
		v.UserID = userID
		v.Status = 1
		if v.ID != 0 {
			tx.Model(&v).Updates(&v)
		} else {
			tx.Create(&v)
		}
	}

	for _, v := range nUser.WorkExps {
		v.UserID = userID
		v.Status = 1
		if v.ID != 0 {
			tx.Model(&v).Updates(&v)
		} else {
			tx.Create(&v)
		}
	}
	err := tx.First(&user, userID).Error
	now := time.Now()
	nUser.UpdatedAt = &now
	err = tx.Model(&user).Updates(nUser).Error
	if err != nil {
		tx.Rollback()
		common.Response(c, "更新失败")
		return
	}
	tx.Commit()
	common.Response(c, user, e.GetMsg(e.SUCCESS), e.SUCCESS)
}

// UpdateInfo 更新用户信息
func UpdateInfo(c iris.Context) {

	var userReqData model.User
	if err := c.ReadJSON(&userReqData); err != nil {
		common.Response(c, "参数无效")
		return
	}
	user := c.Values().Get("user").(*User)

	field := c.FormValue("field")
	resData := make(map[string]interface{})
	resData["id"] = user.ID

	switch field {
	case "sex":
		if userReqData.Sex != model.UserSexMale && userReqData.Sex != model.UserSexFemale {
			common.Response(c, "无效的性别")
			return
		}
		if err := initialize.DB.Model(&user).Update("sex", userReqData.Sex).Error; err != nil {
			common.Response(c, "修改失败")
			return
		}
		resData[field] = userReqData.Sex
	case "signature":
		userReqData.Signature = utils.AvoidXSS(userReqData.Signature)
		userReqData.Signature = strings.TrimSpace(userReqData.Signature)
		// 个性签名可以为空
		if utf8.RuneCountInString(userReqData.Signature) > model.MaxSignatureLen {
			common.Response(c, "个性签名不能超过"+strconv.Itoa(model.MaxSignatureLen)+"个字符")
			return
		}
		if err := initialize.DB.Model(&user).Update("signature", userReqData.Signature).Error; err != nil {
			golog.Error(err)
			common.Response(c, "error")
			return
		}
		resData[field] = userReqData.Signature
	case "location":
		userReqData.Location = utils.AvoidXSS(userReqData.Location)
		userReqData.Location = strings.TrimSpace(userReqData.Location)
		// 居住地可以为空
		if utf8.RuneCountInString(userReqData.Location) > model.MaxLocationLen {
			common.Response(c, "居住地不能超过"+strconv.Itoa(model.MaxLocationLen)+"个字符")
			return
		}
		if err := initialize.DB.Model(&user).Update("location", userReqData.Location).Error; err != nil {
			common.Response(c, "error")
			return
		}
		resData[field] = userReqData.Location
	case "introduce":
		userReqData.Introduction = utils.AvoidXSS(userReqData.Introduction)
		userReqData.Introduction = strings.TrimSpace(userReqData.Introduction)
		// 个人简介可以为空
		if utf8.RuneCountInString(userReqData.Introduction) > model.MaxIntroduceLen {
			common.Response(c, "个人简介不能超过"+strconv.Itoa(model.MaxIntroduceLen)+"个字符")
			return
		}
		if err := initialize.DB.Model(&user).Update("introduce", userReqData.Introduction).Error; err != nil {
			common.Response(c, "error")
			return
		}
		resData[field] = userReqData.Introduction
	default:
		common.Response(c, "参数无效")
		return
	}
	common.Response(c, iris.Map{"data": resData})
}

// UpdatePassword 更新用户密码
func UpdatePassword(c iris.Context) {

	type userReqData struct {
		Password string `json:"password" binding:"required,min=6,max=20"`
		NewPwd   string `json:"newPwd" binding:"required,min=6,max=20"`
	}
	var userData userReqData
	if err := c.ReadJSON(&userData); err != nil {
		common.Response(c, "参数无效")
		return
	}

	userID := c.Values().Get("userID").(uint64)
	var user model.User

	if err := initialize.DB.First(&user, userID).Error; err != nil {
		common.Response(c, "error")
		return
	}

	if checkPassword(userData.Password, &user) {
		user.Password = encryptPassword(userData.NewPwd, userData.NewPwd)
		if err := initialize.DB.Save(&user).Error; err != nil {
			common.Response(c, "原密码不正确")
			return
		}
		common.Response(c, "更新成功")
	} else {
		common.Response(c, "原密码错误")
		return
	}
}

// PublicInfo 用户公开的信息
func PublicInfo(c iris.Context) {

	var userID string

	if userID = c.URLParam("id"); userID != "" {
		common.Response(c, "无效的ID")
		return
	}
	var user model.User
	if err := initialize.DB.First(&user, userID).Error; err != nil {
		common.Response(c, "无效的ID")
		return
	}
	if user.Sex == model.UserSexFemale {
		user.CoverURL = "https://www.golang123.com/upload/img/2017/09/13/d20f62c6-bd11-4739-b79b-48c9fcbce392.jpg"
	} else {
		user.CoverURL = "https://www.golang123.com/upload/img/2017/09/13/e672995e-7a39-4a05-9673-8802b1865c46.jpg"
	}
	common.Response(c, iris.Map{"user": user})
}

// SecretInfo 返回用户信息，包含一些私密字段
func SecretInfo(c iris.Context) {
	if user := c.Values().Get("user"); user != nil {
		common.Response(c,
			iris.Map{
				"user": user,
			}, "success")
	}
}

// InfoDetail 返回用户详情信息(教育经历、职业经历等)，包含一些私密字段
func InfoDetail(c iris.Context) {

	userID := c.Values().Get("userID").(uint64)
	var user model.User

	if err := initialize.DB.First(&user, userID).Error; err != nil {
		common.Response(c, "error")
		return
	}

	if err := initialize.DB.Model(&user).Related(&user.EduExps).Error; err != nil {
		common.Response(c, "error")
		return
	}

	if err := initialize.DB.Model(&user).Related(&user.WorkExps).Error; err != nil {
		common.Response(c, "error")
		return
	}

	if user.Sex == model.UserSexFemale {
		user.CoverURL = "https://www.golang123.com/upload/img/2017/09/13/d20f62c6-bd11-4739-b79b-48c9fcbce392.jpg"
	} else {
		user.CoverURL = "https://www.golang123.com/upload/img/2017/09/13/e672995e-7a39-4a05-9673-8802b1865c46.jpg"
	}

	common.Response(c,
		iris.Map{
			"user": user,
		})
}

// AllList 查询用户列表，只有管理员才能调此接口
func AllList(c iris.Context) {

	user := c.Values().Get("user").(*User)

	allUserRole := []uint8{
		model.UserRoleNormal,
		model.UserRoleEditor,
		model.UserRoleAdmin,
		model.UserRoleCrawler,
		model.UserRoleSuperAdmin,
	}
	foundRole := false
	for _, r := range allUserRole {
		if r == user.Role {
			foundRole = true
			break
		}
	}

	var startTime string
	var endTime string

	if startAt, err := strconv.Atoi(c.FormValue("startAt")); err != nil {
		startTime = time.Unix(0, 0).Format("2006-01-02 15:04:05")
	} else {
		startTime = time.Unix(int64(startAt/1000), 0).Format("2006-01-02 15:04:05")
	}

	if endAt, err := strconv.Atoi(c.FormValue("endAt")); err != nil {
		endTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		endTime = time.Unix(int64(endAt/1000), 0).Format("2006-01-02 15:04:05")
	}

	pageNo, pageNoErr := strconv.Atoi(c.FormValue("pageNo"))
	if pageNoErr != nil {
		pageNo = 1
	}
	if pageNo < 1 {
		pageNo = 1
	}

	offset := (pageNo - 1) * model.PageSize
	pageSize := model.PageSize

	var users []model.User
	var totalCount int
	if foundRole {
		if err := initialize.DB.Model(&model.User{}).Where("created_at >= ? AND created_at < ? AND role = ?", startTime, endTime, user.Role).
			Count(&totalCount).Error; err != nil {
			golog.Error(err)
			common.Response(c, "error")
			return
		}
		if err := initialize.DB.Where("created_at >= ? AND created_at < ? AND role = ?", startTime, endTime, user.Role).
			Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
			golog.Error(err)
			common.Response(c, "error")
			return
		}
	} else {
		if err := initialize.DB.Model(&model.User{}).Where("created_at >= ? AND created_at < ?", startTime, endTime).
			Count(&totalCount).Error; err != nil {
			golog.Error(err)
			common.Response(c, "error")
			return
		}
		if err := initialize.DB.Where("created_at >= ? AND created_at < ?", startTime, endTime).Order("created_at DESC").Offset(offset).
			Limit(pageSize).Find(&users).Error; err != nil {
			golog.Error(err)
			common.Response(c, "error")
			return
		}
	}
	var results []interface{}
	for i := 0; i < len(users); i++ {
		results = append(results, iris.Map{
			"id":          users[i].ID,
			"name":        users[i].Name,
			"email":       users[i].Email,
			"role":        users[i].Role,
			"status":      users[i].Status,
			"createdAt":   users[i].CreatedAt,
			"activatedAt": users[i].ActivatedAt,
		})
	}
	common.Res(c, iris.Map{
		"errNo": e.SUCCESS,
		"msg":   "success",
		"data": iris.Map{
			"users":      results,
			"pageNo":     pageNo,
			"pageSize":   pageSize,
			"totalCount": totalCount,
		},
	})
}

func topN(c iris.Context, n int) {

	var users []model.User
	if err := initialize.DB.Order("score DESC").Limit(n).Find(&users).Error; err != nil {
		golog.Error(err)
		common.Response(c, "error")
	} else {
		common.Response(c,
			iris.Map{
				"users": users,
			})
	}
}

// Top10 返回积分排名前10的用户
func Top10(c iris.Context) {
	topN(c, 10)
}

// Top100 返回积分排名前100的用户
func Top100(c iris.Context) {
	topN(c, 100)
}

// UploadAvatar 上传用户头像
func UploadAvatar(c iris.Context) {

	data := upload.Upload(c)
	if data == nil {
		return
	}

	avatarURL := data.URL
	user := c.Values().Get("user").(*User)

	if err := initialize.DB.Model(&user).Update("avatar_url", avatarURL).Error; err != nil {
		return
	}

	common.Response(c, data)
}

// AddCareer 添加职业经历
func AddCareer(c iris.Context) {

	var work model.Work
	if err := c.ReadJSON(&work); err != nil {
		common.Response(c, "参数无效")
		return
	}

	work.Company = strings.TrimSpace(utils.AvoidXSS(work.Company))
	work.Title = utils.AvoidXSS(work.Title)
	work.Title = strings.TrimSpace(work.Title)

	if work.Company == "" {
		common.Response(c, "公司或组织名称不能为空")
		return
	}

	if utf8.RuneCountInString(work.Company) > model.MaxCareerCompanyLen {
		common.Response(c, "公司或组织名称不能超过"+strconv.Itoa(model.MaxCareerCompanyLen)+"个字符")
		return
	}

	if work.Title == "" {
		common.Response(c, "职位不能为空")
		return
	}

	if utf8.RuneCountInString(work.Title) > model.MaxCareerTitleLen {
		common.Response(c, "职位不能超过"+strconv.Itoa(model.MaxCareerTitleLen)+"个字符")
		return
	}

	user := c.Values().Get("user").(*User)

	work.UserID = user.ID
	work.CreatedAt = time.Now()

	if err := initialize.DB.Create(&work).Error; err != nil {
		common.Response(c, "error")
		return
	}

	common.Response(c, work)
}

// AddSchool 添加教育经历
func AddSchool(c iris.Context) {

	var edu model.Education
	if err := c.ReadJSON(&edu); err != nil {
		common.Response(c, "参数无效")
		return
	}

	edu.School = strings.TrimSpace(utils.AvoidXSS(edu.School))
	edu.Speciality = utils.AvoidXSS(edu.Speciality)
	edu.Speciality = strings.TrimSpace(edu.Speciality)

	if edu.School == "" {
		common.Response(c, "学校或教育机构名不能为空")
		return
	}

	if utf8.RuneCountInString(edu.School) > model.MaxSchoolNameLen {
		common.Response(c, "学校或教育机构名不能超过"+strconv.Itoa(model.MaxSchoolNameLen)+"个字符")
		return
	}

	if edu.Speciality == "" {
		common.Response(c, "专业方向不能为空")
		return
	}

	if utf8.RuneCountInString(edu.Speciality) > model.MaxSchoolSpecialityLen {
		common.Response(c, "专业方向不能超过"+strconv.Itoa(model.MaxSchoolSpecialityLen)+"个字符")
		return
	}

	user := c.Values().Get("user").(*User)

	edu.UserID = user.ID
	edu.CreatedAt = time.Now()

	if err := initialize.DB.Create(&edu).Error; err != nil {
		common.Response(c, "error")
		return
	}

	common.Response(c, edu)
}

//删除职业经历
func DeleteWork(c iris.Context) {

	id := c.Params().GetUint64Default("id", 0)

	var work model.Work
	if err := initialize.DB.First(&work, id).Error; err != nil {
		common.Response(c, "无效的id.")
		return
	}

	if err := initialize.DB.Delete(&work).Error; err != nil {
		common.Response(c, "error")
		return
	}

	common.Response(c, iris.Map{"id": work.ID})

}

// 删除教育经历
func DeleteEdu(c iris.Context) {

	id := c.Params().GetUint64Default("id", 0)

	var school model.School
	if err := initialize.DB.First(&school, id).Error; err != nil {
		common.Response(c, "无效的id.")
		return
	}

	if err := initialize.DB.Delete(&school).Error; err != nil {
		common.Response(c, "error")
		return
	}
	common.Response(c, iris.Map{"id": school.ID})
}
func CheckAuth(username, password string) (bool, error) {
	var auth model.Auth
	err := initialize.DB.Select("id").Where(model.Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

// UserFromRedis 从redis中取出用户信息
func UserFromRedis(userID uint64) (*User, error) {
	loginUser := model.LoginUser + strconv.FormatUint(userID, 10)

	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("SELECT", credis.UserIndex)
	userString, err := redis.String(conn.Do("GET", loginUser))
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	var user User
	err = utils.Json.UnmarshalFromString(userString, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserLastActiveTime(userID uint64) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	err := conn.Send("SELECT", credis.CronIndex)
	_, err = conn.Do("ZADD", model.LoginUser+"ActiveTime",
		time.Now().Unix(), strconv.FormatUint(userID, 10))
	if err != nil {
		golog.Error(err)
		return err
	}
	return nil
}

// UserToRedis 将用户信息存到redis
func UserToRedis(user *User) error {
	UserString, err := utils.Json.MarshalToString(user)
	if err != nil {
		return err
	}
	loginUserKey := model.LoginUser + strconv.FormatUint(user.ID, 10)

	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("SELECT", credis.UserIndex)
	if _, redisErr := conn.Do("SET", loginUserKey, UserString, "EX", initialize.Config.Server.TokenMaxAge); redisErr != nil {
		return errors.New("error")
	}
	return nil
}

func EditUserRedis(user *User) error {
	UserString, err := utils.Json.MarshalToString(user)
	if err != nil {
		return err
	}
	loginUserKey := model.LoginUser + strconv.FormatUint(user.ID, 10)

	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("SELECT", credis.UserIndex)
	if _, redisErr := conn.Do("SET", loginUserKey, UserString); redisErr != nil {
		return errors.New("error")
	}
	return nil
}

// CheckPassword 验证密码是否正确
func checkPassword(password string, user *model.User) bool {
	if password == "" || user.Password == "" {
		return false
	}
	return encryptPassword(password, password) == user.Password
}

// Salt 每个用户都有一个不同的盐
func salt(password string) string {

	return password[0:5]
}

// EncryptPassword 给密码加密
func encryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum(utils.ToBytes(password)))
	hash = salt + password + initialize.Config.Server.PassSalt
	hash = fmt.Sprintf("%x", md5.Sum(utils.ToBytes(hash)))
	return
}

type Claims struct {
	UserID   uint64 `json:"user_id"`
	UserRole uint8  `json:"user_role"`
	jwt.StandardClaims
}

func GenerateToken(user *model.User) (string, error) {

	claims := Claims{
		UserID:   user.ID,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + initialize.Config.Server.TokenMaxAge,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hoper",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(utils.ToBytes(initialize.Config.Server.JwtSecret))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, _ := (&jwt.Parser{SkipClaimsValidation:true}).ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return utils.ToBytes(initialize.Config.Server.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			now:=time.Now().Unix()
			if claims.VerifyExpiresAt(now, false) == false {
				return nil,errors.New("登录超时")
			}
			return claims, nil
		}
	}

	return nil, errors.New("未登录")
}
