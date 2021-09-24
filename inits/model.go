package inits

type (
	Conf struct {
		DbConf    *DbConf
		GinConf   *GinConf
		RedisConf *RedisConf
		MqConf    *MqConf
	}

	DbConf struct {
		Host     string `toml:"host"`
		DbName   string `toml:"dbname"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Sslmode  string `toml:"disable"`
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
		Db        string `toml:"db"`
	}

	MqConf struct {
		Server []string `toml:"server"`
	}
)
