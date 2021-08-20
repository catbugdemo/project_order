package inits

var (
	mqConfig *MqConfig
)

type MqConfig struct {
	Server []string `toml:"server"`
}

func init() {
	InitMqConfig()
	InitRocketMq()
}

func InitMqConfig() {
	c := GetConfig()
	mqConfig = &MqConfig{
		Server: c.GetStringSlice("server"),
	}
}

func InitRocketMq()  {

}
