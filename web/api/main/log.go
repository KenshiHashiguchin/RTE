package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func logSetting(r *gin.Engine) {
	// ファイルへログを書き込む
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// ログに書き込みつつ、コンソールにも出力する場合、下記のコードを利用する。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
}
