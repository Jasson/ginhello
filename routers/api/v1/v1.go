package v1

import (

	"github.com/gin-gonic/gin"
	"octopus_ad/routers/api/v1/user"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		// 用户路由
		user.RegisterRouter(v1.Group("/user"))
	}
}
