package setting

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	Port string
}

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
}

var AppConfig = &App{}

func SetUp() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置失败: %s \n", err)
	}
	AppConfig.Port = viper.GetString("app.port")
}
