package inits

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

var (
	ginConfig *GinConfig
)

type GinConfig struct {
	Port         string `toml:"port"`
	ReadTimeout  int    `toml:"read_timeout"`
	WriteTimeout int    `toml:"write_timeout"`
}

func InitHttp() {
	InitGinConfig()
	InitGin()
}

func InitGinConfig() {
	c := GetConfig()
	ginConfig = &GinConfig{
		Port:         c.GetString("http.port"),
		ReadTimeout:  c.GetInt("http.read_timeout"),
		WriteTimeout: c.GetInt("http.write_timeout"),
	}
}

func InitGin() {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/version/", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": "v1.0.0"})
	})

	s := &http.Server{
		Addr:         ginConfig.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(ginConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(ginConfig.WriteTimeout) * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("http: Server Close:%+v", errors.WithStack(err))
		}
		log.Fatalf("http开启监听服务失败:%+v", errors.WithStack(err))
	}

}
