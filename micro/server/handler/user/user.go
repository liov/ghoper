package user

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"hoper/client/controller/mail"
	"hoper/initialize"
	"hoper/model"
	"hoper/protobuf"
	"hoper/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	activeDuration = 24 * 60 * 60
	resetDuration  = 24 * 60 * 60
)

type UserHandler struct {
}

func (u *UserHandler) Signup(ctx context.Context, signupReq *protobuf.SignupReq, loginRep *protobuf.LoginRep) error {

	signupReq.Name = utils.AvoidXSS(signupReq.Name)
	signupReq.Name = strings.TrimSpace(signupReq.Name)
	signupReq.Email = strings.TrimSpace(signupReq.Email)

	if strings.Index(signupReq.Name, "@") != -1 {
		return errors.New("用户名中不能含有@字符")
	}

	var user model.User
	if err := initialize.DB.Where("email = ? OR phone = ?", signupReq.Email, signupReq.Phone).Find(&user).Error; err == nil {
		if user.Phone != nil && *user.Phone == signupReq.Phone {
			return errors.New("手机号已被注册")
		} else if user.Email == signupReq.Email {
			return errors.New("邮箱已存在")
		}
	}

	var newUser protobuf.User
	nowTime := time.Now()
	newUser.Name = signupReq.Name
	newUser.Email = signupReq.Email
	newUser.Phone = signupReq.Phone
	newUser.Password = encryptPassword(signupReq.Password, signupReq.Password)
	//newUser.Role = model.UserRoleNormal
	newUser.Status = model.UserStatusInActive

	if err := initialize.DB.Create(&newUser).Error; err != nil {
		return errors.New("error")
	}

	activeUser := model.ActiveTime + strconv.FormatUint((uint64)(newUser.ID), 10)

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	currTime := nowTime.Unix()

	if _, err := RedisConn.Do("SET", activeUser, currTime, "EX", activeDuration); err != nil {
	}

	go func() {
		sendMail("/active", "账号激活", currTime, newUser)
	}()

	loginRep.Msg = "注册成功"
	loginRep.User = &newUser

	return nil
}

func (u *UserHandler) Login(ctx context.Context, loginReq *protobuf.LoginReq, loginRep *protobuf.LoginRep) error {

	var signinInput, password, luosimaoRes, sql string

	if loginReq.Input != "" {
		emailMatch, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+.[a-zA-Z]{2,3}$`, loginReq.Input)

		phoneMatch, _ := regexp.MatchString(`^1[0-9]{10}$`, loginReq.Input)
		if emailMatch {
			sql = "email = ?"
		} else if phoneMatch {
			sql = "phone = ?"
		} else {
			return errors.New("账号错误")
		}
	}

	signinInput = loginReq.Input
	password = loginReq.Password
	luosimaoRes = loginReq.LuosimaoRes

	verifyErr := utils.LuosimaoVerify(initialize.Config.Server.LuosimaoVerifyURL, initialize.Config.Server.LuosimaoAPIKey, luosimaoRes)

	if verifyErr != nil {
		loginRep.Msg = verifyErr.Error()
		return nil
	}

	var user protobuf.User
	if err := initialize.DB.Where(sql, signinInput).Find(&user).Error; err != nil {
		return errors.New("账号不存在")
	}

	if checkPassword(password, user) {
		if user.Status == model.UserStatusInActive {
			encodedEmail := base64.StdEncoding.EncodeToString(utils.ToBytes(user.Email))
			loginRep.Msg = "账号未激活"
			loginRep.User.Email = encodedEmail
			return nil
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID,
		})
		tokenString, err := token.SignedString(utils.ToBytes(initialize.Config.Server.TokenSecret))
		if err != nil {
			return errors.New("内部错误")
		}

		if err := UserToRedis(user); err != nil {
			return errors.New("内部错误")
		}

		loginRep.Msg = "登录成功"
		loginRep.Token = tokenString
		loginRep.User = &user

		/*		session := sessions.Default(c)
				session.Set("user", user)
				session.Save()*/
		//userBytes, err := json.Marshal(user)
		//c.SetCookie("user", string(userBytes), initialize.ServerSettings.TokenMaxAge, "/", "hoper.xyz", false, true)

		return nil
	}

	return errors.New("账号或密码错误")
}

func (u *UserHandler) Logout(ctx context.Context, logoutReq *protobuf.LogoutReq, logoutRep *protobuf.LogoutRep) error {

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("DEL", model.LoginUser+strconv.FormatUint(logoutReq.ID, 10)); err != nil {
		return err
	}

	logoutRep.Msg = "已退出"

	return nil
}

func (u *UserHandler) GetUser(ctx context.Context, getReq *protobuf.GetReq, user *protobuf.User) error {
	initialize.DB.Where("id = ?", getReq.ID).Find(&user)
	return nil
}

// EncryptPassword 给密码加密
func encryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum(utils.ToBytes(password)))
	hash = salt + password + initialize.Config.Server.PassSalt
	hash = fmt.Sprintf("%x", md5.Sum(utils.ToBytes(hash)))
	return
}

func sendMail(action string, title string, curTime int64, user protobuf.User) {
	siteName := initialize.Config.Server.SiteName
	siteURL := "http://" + initialize.Config.Server.Host
	secretStr := strconv.Itoa((int)(curTime)) + user.Email + user.Password
	secretStr = fmt.Sprintf("%x", md5.Sum(utils.ToBytes(secretStr)))
	actionURL := siteURL + "/user" + action + "/"

	actionURL = actionURL + strconv.FormatUint((uint64)(user.ID), 10) + "/" + secretStr
	fmt.Println(actionURL)
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

// UserFromRedis 从redis中取出用户信息
func UserFromRedis(userID int) (protobuf.User, error) {
	loginUser := model.LoginUser + strconv.Itoa(userID)

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	userBytes, err := redis.String(RedisConn.Do("GET", loginUser))
	if err != nil {
		fmt.Println(err)
		return protobuf.User{}, errors.New("未登录")
	}
	var user protobuf.User
	bytesErr := utils.Json.UnmarshalFromString(userBytes, &user)
	if bytesErr != nil {
		fmt.Println(bytesErr)
		return user, errors.New("未登录")
	}
	return user, nil
}

// UserToRedis 将用户信息存到redis
func UserToRedis(user protobuf.User) error {
	UserString, err := utils.Json.MarshalToString(user)
	if err != nil {
		return errors.New("error")
	}
	loginUserKey := model.LoginUser + strconv.FormatUint((uint64)(user.ID), 10)

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, redisErr := RedisConn.Do("SET", loginUserKey, UserString, "EX", initialize.Config.Server.TokenMaxAge); redisErr != nil {
		return errors.New("缓存用户信息出错")
	}
	return nil
}

// CheckPassword 验证密码是否正确
func checkPassword(password string, user protobuf.User) bool {
	if password == "" || user.Password == "" {
		return false
	}
	return encryptPassword(password, password) == user.Password
}
