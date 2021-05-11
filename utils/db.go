package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GetDb 获取数据库连接
func GetDb() (*gorm.DB, func(), error) {

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	// clean db connection for defer
	cleanFunc := func() {
		sqlDb, err := db.DB()
		if sqlDb != nil {
			_ = sqlDb.Close()
		}
		if err != nil {
			panic(err)
		}
	}

	return db, cleanFunc, nil
}
