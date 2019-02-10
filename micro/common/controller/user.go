package controller

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"micro/common/controller/common"
	"micro/common/initialize"
	"micro/common/model"
	"micro/protobuf"
	"strconv"
)

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
	bytesErr := common.Json.UnmarshalFromString(userBytes, &user)
	if bytesErr != nil {
		fmt.Println(bytesErr)
		return user, errors.New("未登录")
	}
	return user, nil
}

// UserToRedis 将用户信息存到redis
func UserToRedis(user protobuf.User) error {
	UserString, err := common.Json.MarshalToString(user)
	if err != nil {
		return errors.New("error")
	}
	loginUserKey := model.LoginUser + strconv.FormatUint((uint64)(user.ID), 10)

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	if _, redisErr := RedisConn.Do("SET", loginUserKey, UserString, "EX", initialize.Config.Server.TokenMaxAge); redisErr != nil {
		return errors.New("error")
	}
	return nil
}
