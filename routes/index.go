package routes

import (
	"fmt"
	"smg/app/controller"
	"time"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
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
	router.Use(gin.Recovery())

	regsiterRouter(engine)
}

// 注册路由
func regsiterRouter(router *gin.Engine) {
	new(controller.UserController).Router(router)
}
