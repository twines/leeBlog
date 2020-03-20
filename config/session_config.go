/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/8 10:17
*/
package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"leeblog.com/pkg/utils"
)

var store sessions.Store

func SessionConfig(keyPairs string) sessions.Store {
	config := utils.Config()
	sessionStore := config.GetString("session.Store")
	sessionMaxAge := config.GetInt("session.MaxAge")
	sessionSecret := config.GetString("session.Secret")
	switch sessionStore {
	case "cookie":
		store = cookie.NewStore([]byte(sessionSecret))
		break
	case "redis":
		redisConfig := RedisConfig()
		store, _ = redis.NewStore(11, "tcp", redisConfig.Addr+":"+redisConfig.Port, redisConfig.Password, []byte(sessionSecret))
		break
	}
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //30min
		Path:   "/",
	})
	return store
}
