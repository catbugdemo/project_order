package inits

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"time"
)

var (
	pool        *redis.Pool
	redisConfig *RedisConfig
)

type RedisConfig struct {
	Host      string `toml:"host"`
	Password  string `toml:"passowrd"`
	MaxIdle   int    `toml:"max_idle"`
	MaxActive int    `toml:"max_active"`
	Db        string `toml:"db"`
}

func init() {
	InitRedisConfig()
	InitRedisPool()
}

// InitRedisConfig 初始化缓存配置
func InitRedisConfig() {
	c := GetConfig()

	redisConfig = &RedisConfig{
		Host:      c.GetString("redis.host"),
		Password:  c.GetString("redis.password"),
		MaxIdle:   c.GetInt("redis.max_idle"),
		MaxActive: c.GetInt("redis.max_active"),
	}
}

// InitRedisPool 初始化缓存池
func InitRedisPool() {
	pool = &redis.Pool{
		//最大闲置连接
		MaxIdle: redisConfig.MaxIdle,
		//最大活动数
		MaxActive: redisConfig.MaxActive,
		//数据库连接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConfig.Host)
			if err != nil {
				c.Close()
				log.Printf("fail to dial redis: %+v\n", errors.WithStack(err))
				return nil, err
			}
			//密码认证
			if redisConfig.Password != "" {
				if _, err = c.Do("AUTH", redisConfig.Password); err != nil {
					c.Close()
					log.Printf("fail to auth redis: %+v\n", errors.WithStack(err))
					return nil, err
				}
			}
			//redis 缓存数据库认证
			if redisConfig.Db != "" {
				if _, err = c.Do("SELECT", redisConfig.Db); err != nil {
					c.Close()
					log.Printf("fail to SELECT DB redis: %+v\n", errors.WithStack(err))
					return nil, err
				}
			}
			return c, err
		},
		//测试连接是否正常
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				c.Close()
				log.Printf("fail to ping redis: %+v\n", err)
				return err
			}
			return nil
		},
	}
}

// Get 获取缓存连接
func Get() redis.Conn {
	return pool.Get()
}
