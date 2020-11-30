package routers

import (
	"main/services/user"

	"github.com/gin-gonic/gin"
)

// Router gin Engine
var Router *gin.Engine

// 初始化
func init() {
	Router = gin.Default()
}

// RegisterRouter 注册路由
func RegisterRouter() {
	v1 := Router.Group("/v1")
	{
		// 客服用户
		{
			// 单条用户
			v1.GET("/user/:id", user.GetUser)
		}
	}

}
