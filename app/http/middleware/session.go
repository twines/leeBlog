package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"leeblog.com/config"
)

func Session(keyPairs string) gin.HandlerFunc {
	store := config.SessionConfig(keyPairs)
	return sessions.Sessions(keyPairs, store)
}
