package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context, key string, model interface{}) {
	session := sessions.Default(c)
	session.Set(key, model)
	_ = session.Save()
}
func AuthCheck(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}
