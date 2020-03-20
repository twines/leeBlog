package config

import "leeblog.com/pkg/utils"

type Server struct {
	Debug bool
	Port  string
}

var (
	server Server
)

func ServerConfig() Server {
	viper := utils.Config()
	if err := viper.ReadInConfig(); err == nil {
		_ = viper.UnmarshalKey("server", &server)
	}
	return server
}
