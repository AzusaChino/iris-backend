package setting

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	Port string
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
}

var AppConfig = &App{}
var MysqlConfig = &Mysql{}

func SetUp() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置失败: %s \n", err)
	}
	AppConfig.Port = viper.GetString("app.port")

	MysqlConfig.Host = viper.GetString("mysql.host")
	MysqlConfig.Port = viper.GetInt("mysql.port")
	MysqlConfig.Username = viper.GetString("mysql.username")
	MysqlConfig.Password = viper.GetString("mysql.password")
	MysqlConfig.Database = viper.GetString("mysql.database")
}
