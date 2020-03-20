/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/8 9:51
*/
package config

import (
	"leeblog.com/pkg/utils"
)

var (
	r Redis
)

type Redis struct {
	Addr     string
	Password string
	DB       int
	Port     string
}

func RedisConfig() Redis {
	viper := utils.Config()
	if err := viper.ReadInConfig(); err == nil {
		_ = viper.UnmarshalKey("redis", &r)
	}
	return r
}
