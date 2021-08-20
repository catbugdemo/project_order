package inits

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/pkg/errors"
	"time"
)

var (
	db       *gorm.DB
	dbConfig *DbConfig
)

type DbConfig struct {
	Host     string `toml:"host"`
	DbName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Sslmode  string `toml:"disable"`
}

func init() {
	InitDbConfig()
	InitDB()
}

func InitDbConfig() {
	c := GetConfig()

	dbConfig = &DbConfig{
		Host:     c.GetString("database.host"),
		DbName:   c.GetString("database.dbname"),
		User:     c.GetString("database.user"),
		Password: c.GetString("database.password"),
		Sslmode:  c.GetString("database.sslmode"),
	}
}

func InitDB() {

	params := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.DbName,
		dbConfig.Sslmode,
		dbConfig.Password,
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

// 获取数据库
func DB() *gorm.DB {
	return db
}
