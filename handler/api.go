package handler

import "github.com/gin-gonic/gin"

func Api(r *gin.Engine) {
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})
}
