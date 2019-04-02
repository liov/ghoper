package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"hoper/initialize"
	"hoper/utils/hlog"
)

func Set(key string, data interface{}, time int) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func SetList(key string, data interface{}, time int) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	/*	rv := reflect.ValueOf(data)
		switch rv.Kind() {
		case reflect.Array:
		case reflect.Ptr:
		}

		var values []interface{}

		if isJson {
			for i, v := range data {
				value, err := json.Marshal(v)
				if err != nil {
					return err
				}
			values[i] = value
				_, err = conn.Do("LPUSH", key, values)
				if err != nil {
					return err
				}
			}
			//values = append(values,values...)
		}*/

	value, err := json.Marshal(data)

	_, err = conn.Do("LPUSH", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func redisMoments(key string, model interface{}) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if exist, err := redis.Bool(conn.Do("EXISTS", key)); exist && err == nil {
		data, err := redis.Bytes(conn.Do("GET", key))
		if err != nil {
			hlog.Info(err)
			return err
		} else {
			json.Unmarshal(data, model)
			/*	for _, mv := range *moments {
						//瞬间是不需要设置缓存的，前端存储
						mkey := strings.Join([]string{
								e.CacheMoment,
								strconv.FormatUint(uint64(mv.ID),10),
							}, "_")

							mv.BrowseCount = mv.BrowseCount + 1

							_, err =conn.Do("SET", mkey, mv)
							_, err =conn.Do("EXPIRE", mkey, 60)

				}
				if err != nil {
					return err
				}
			*/
		}
	}
	return nil
}
