package middleware

import (
	"fmt"
	"gin_mal_tmp/pkg/e"
	"gin_mal_tmp/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.Success
		token := ctx.GetHeader("access_token")
		if token == "" {
			panic("fuck JWT")
			code = http.StatusNotFound
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				fmt.Println(claims, "|||||", err)
				panic("fuck JWT2")
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.Success {
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
