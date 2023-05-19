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

var runViper *viper.Viper

func init() {
	runViper = viper.New()
	runViper.SetConfigName("app")
	runViper.SetConfigType("toml")
	runViper.AddConfigPath(".")
}

var AppConfig = &App{}
var MysqlConfig = &Mysql{}

func SetUp() {
	err := runViper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置失败: %s \n", err)
	}
	AppConfig.Port = runViper.GetString("app.port")

	MysqlConfig.Host = runViper.GetString("mysql.host")
	MysqlConfig.Port = runViper.GetInt("mysql.port")
	MysqlConfig.Username = runViper.GetString("mysql.username")
	MysqlConfig.Password = runViper.GetString("mysql.password")
	MysqlConfig.Database = runViper.GetString("mysql.database")
}
