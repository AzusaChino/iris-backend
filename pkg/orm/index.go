package orm

import (
	"fmt"
	"iris/pkg/setting"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type ZapLogger struct{}

var db *gorm.DB

func SetUp() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.MysqlConfig.Username,
			setting.MysqlConfig.Password,
			setting.MysqlConfig.Host,
			setting.MysqlConfig.Port,
			setting.MysqlConfig.Database),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "tb_",
		},
		PrepareStmt: true,
		NowFunc:     time.Now().Local,
		Logger:      defaultLogger(),
	})

	if err != nil {
		log.Fatalf("db setup error: %v", err)
	}
}

func GetDb() *gorm.DB {
	return db
}

func defaultLogger() logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Info,
		Colorful:      true,
	})
}
