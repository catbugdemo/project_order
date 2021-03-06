package inits

import (
	"fmt"
	"github.com/catbugdemo/project_order/log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver

	"github.com/pkg/errors"
)

var (
	db *gorm.DB
)

func InitDB() {
	dbConfig := conf.DbConf

	params := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.DbName,
		dbConfig.Password,
		dbConfig.Sslmode,
	)

	open, err := gorm.Open("postgres", params)
	if err != nil {
		log.Fatalf("连接数据库失败:%+v", errors.WithStack(err))
	}
	if err = open.DB().Ping(); err != nil {
		log.Fatalf("ping 数据库失败:%+v", errors.WithStack(err))
	}

	// 设置表不为复数
	open.SingularTable(true)
	// 设置打印日志
	open.LogMode(true)
	// 设置可重复使用连接的最长时间
	open.DB().SetConnMaxLifetime(10 * time.Second)

	db = open
}

// DB 获取数据库
func DB() *gorm.DB {
	return db
}
