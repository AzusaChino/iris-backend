package orm

import (
	"fmt"
	"iris/pkg/setting"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	})

	if err != nil {
		log.Fatalf("db setup error: %v", err)
	}
}

func GetDb() *gorm.DB {
	return db
}
