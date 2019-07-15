package utils

import (
	"octopus_ad/common/pkg/e"
	"time"

	"github.com/gin-gonic/gin"
)

// RES 返回信息自动根据code插入message
func RES(c *gin.Context, code int, obj gin.H) {
	if obj["message"] == "" {
		obj["message"] = e.GetMessage(code)
	}
	obj["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	obj["code"] = code
	c.JSON(code, obj)
}
