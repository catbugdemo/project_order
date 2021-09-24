package inits

import (
	"sync"

	"github.com/catbugdemo/project_order/utils/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	conf   Conf

	log = logger.New()

	mu         sync.RWMutex
	onceConf   sync.Once
	onceStruct sync.Once
)

func init() {
	onceConf.Do(InitConfig)
	onceStruct.Do(InitStructConfig)

	InitDB()
	InitRedisPool()
	InitRocketMq()
}

// InitConfig 初始化配置
func InitConfig() {
	v := viper.New()
	// 初始化配置信息
	v.SetConfigName("local")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	// TODO 项目名称进行配置
	v.AddConfigPath("$GOPATH/src/github.com/catbugdemo/project_order/configs")
	v.AddConfigPath("../configs")

	if err := ReadInConfig(v); err != nil {
		log.Fatalf("读取配置失败：%+v\n", err)
	}
	// 当配置改变是重新运行
	v.WatchConfig()

	config = v
}

// ReadInConfig 通过读写锁保证内容稳定内容
func ReadInConfig(v *viper.Viper) error {
	mu.RLock()
	defer mu.RUnlock()
	if err := v.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InitStructConfig 初始化结构体
func InitStructConfig() {
	conf.DbConf = &DbConf{
		Host:     config.GetString("database.host"),
		DbName:   config.GetString("database.dbname"),
		User:     config.GetString("database.user"),
		Password: config.GetString("database.password"),
		Sslmode:  config.GetString("database.sslmode"),
	}

	conf.GinConf = &GinConf{
		Port:         config.GetString("http.port"),
		ReadTimeout:  config.GetInt("http.read_timeout"),
		WriteTimeout: config.GetInt("http.write_timeout"),
	}

	conf.RedisConf = &RedisConf{
		Host:      config.GetString("redis.host"),
		Password:  config.GetString("redis.password"),
		MaxIdle:   config.GetInt("redis.max_idle"),
		MaxActive: config.GetInt("redis.max_active"),
	}

	conf.MqConf = &MqConf{
		Server: config.GetStringSlice("server"),
	}

}
