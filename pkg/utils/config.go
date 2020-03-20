/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/3 16:50
*/
package utils

import (
	"github.com/spf13/viper"
	"os"
)

var (
	v *viper.Viper
)

func init() {
	v = viper.New()
}
func Config() *viper.Viper {
	if dir, err := os.Getwd(); err == nil {
		if _, err := os.Stat(dir + "/env.example.yml"); err == nil {
			v.SetConfigName("env.example") //设置配置文件的名字
		} else {
			v.SetConfigName("env") //设置配置文件的名字
		}
		v.AddConfigPath(dir)    //添加配置文件所在的路径
		v.SetConfigType("yaml") // or viper.SetConfigType("YAML")
		_ = v.ReadInConfig()
	}
	return v
}
