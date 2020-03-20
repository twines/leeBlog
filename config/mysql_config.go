/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/8 9:39
*/
package config

import (
	"leeblog.com/pkg/utils"
)

var (
	db MySql
)

type MySql struct {
	Port     int
	Host     string
	Username string
	Password string
	DbName   string
	Charset  string
}

func MySQLConfig() MySql {
	viper := utils.Config()
	if err := viper.ReadInConfig(); err == nil {
		_ = viper.UnmarshalKey("mysql", &db)
	}
	return db
}
