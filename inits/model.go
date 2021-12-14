package inits

type (
	Conf struct {
		DbConf    DbConf    `toml:"database"`
		GinConf   GinConf   `toml:"http"`
		RedisConf RedisConf `toml:"redis"`
		MqConf    MqConf    `toml:"rocketmq"`
	}

	DbConf struct {
		Host     string `toml:"host"`
		DbName   string `toml:"dbname"`
		User     string `toml:"user"`
		Port     int    `toml:"port"`
		Password string `toml:"password"`
		Sslmode  string `toml:"sslmode"`
	}

	GinConf struct {
		Port         string `toml:"port"`
		ReadTimeout  int    `toml:"read_timeout"`
		WriteTimeout int    `toml:"write_timeout"`
	}

	RedisConf struct {
		Host      string `toml:"host"`
		Password  string `toml:"passowrd"`
		MaxIdle   int    `toml:"max_idle"`
		MaxActive int    `toml:"max_active"`
		Db        int    `toml:"db"`
	}

	MqConf struct {
		Server []string `toml:"server"`
	}
)
