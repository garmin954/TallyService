package controller

import (
	"github.com/gin-gonic/gin"
)

type Common struct{}

// 路由
func (common *Common) Router(engine *gin.RouterGroup) {
	// group := engine.Group("/c")
	// route := group.Use(middleware.JwtToken())
	// route.GET("/query_users", user.queryUsers)

	// engine.POST("/login", user.login)
	// engine.POST("/register", user.register)
}
