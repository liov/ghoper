package controller

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/sessions"
	"time"
)

func init() {
	//github.com/gin-contrib/sessions 获取User必须gob注册
	//必须encoding/gob编码解码进行注册
	//gob.Register(&User{})

	// AES仅支持16,24或32字节的密钥大小。
	//您需要准确提供该密钥字节大小，或者从您键入的内容中获取密钥。
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie = securecookie.New(hashKey, blockKey)

	Sess = sessions.New(sessions.Config{
		Cookie: "hopersid",
		//Encode:       secureCookie.Encode,
		//Decode:       secureCookie.Decode,
		Expires:      45 * time.Minute, // <=0 means unlimited life. Defaults to 0.
		AllowReclaim: true,
	})

	//Sess.UseDatabase(initialize.BoltDB)

	//Gsess = memstore.NewMemStore(hashKey, blockKey)
}
