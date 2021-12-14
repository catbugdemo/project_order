package inits

import (
	"github.com/catbugdemo/project_order/handler"
	"github.com/catbugdemo/project_order/log"
	"github.com/catbugdemo/project_order/middleware"
	"github.com/gin-contrib/cors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func InitGin() {
	ginConfig := conf.GinConf

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// 跨域设置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"POST", "PUT", "GET", "DELETE", "PATCH"}
	corsConfig.AllowHeaders = []string{"Origin", "token", "Content-Type", "Accept", "Access-Control-Allow-Headers"}
	router.Use(cors.New(corsConfig))

	// 时间测试
	router.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{"time": time.Now()})
	})

	r := router.Use(middleware.Auth("jwtSecret"))

	// 登录
	{
		router.POST("/login/wechat", handler.WechatLogin)
	}

	// 路由
	{
		r.POST("")
	}

	s := &http.Server{
		Addr:         ginConfig.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(ginConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(ginConfig.WriteTimeout) * time.Second,
	}

	log.Infof("run at %v\n", ginConfig.Port)
	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Infof("http: Server Close:%+v\n", errors.WithStack(err))
		}
		log.Fatalf("http开启监听服务失败:%+v", errors.WithStack(err))
	}

}
