package inits

import (
	"net/http"
	"time"

	"github.com/catbugdemo/project_order/handler"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func InitGin() {
	ginConfig := conf.GinConf

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/version/", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": "v1.0.0"})
	})

	handler.Api(router)

	s := &http.Server{
		Addr:         ginConfig.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(ginConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(ginConfig.WriteTimeout) * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("http: Server Close:%+v\n", errors.WithStack(err))
		}
		log.Fatalf("http开启监听服务失败:%+v", errors.WithStack(err))
	}

}
