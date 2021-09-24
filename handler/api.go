package handler

import "github.com/gin-gonic/gin"

// API 测试接口
func API(r *gin.Engine) {
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})
}
