package inits

import (
	"github.com/BurntSushi/toml"
	"log"
	"sync"
)

var (
	conf     Conf
	onceConf sync.Once
)

func init() {
	onceConf.Do(InitConfig)
	InitDB()
	InitRedisPool()
	InitRocketMq()
}

// InitConfig 初始化配置
func InitConfig() {
	// TODO 修改为自己的配置文件路径
	if _, err := toml.DecodeFile("/Users/zonst/go/src/github.com/catbugdemo/project_order/docs/local.toml", &conf); err != nil {
		log.Fatalf("读取配置失败：%+v\n", err)
	}
}
