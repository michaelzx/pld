package pld_db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/michaelzx/pld/pld_config"
	"github.com/michaelzx/pld/pld_logger"
	"log"
)

var db *gorm.DB

func InitDB(appDbCfg *pld_config.DbConfig) *gorm.DB {
	// loc=Local,标识跟随系统
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", // &charset=utf8
		appDbCfg.Usr,
		appDbCfg.Psw,
		appDbCfg.Host,
		appDbCfg.Port,
		appDbCfg.DbName,
	)
	var err error
	if db, err = gorm.Open("mysql", connStr); err == nil {
		// 连接最长存活期，超过这个时间连接将不再被复用
		// db.DB().SetConnMaxLifetime(1 * time.Second)
		// 最大空闲连接数
		// db.DB().SetMaxIdleConns(-1)
		// 数据库最大连接数
		// db.DB().SetMaxOpenConns(120)

		db.SingularTable(true)
		if appDbCfg.MaxLifetime > 0 {
			db.DB().SetConnMaxLifetime(appDbCfg.MaxLifetime)
		}
		if appDbCfg.MaxIdleConns > 0 {
			db.DB().SetMaxIdleConns(appDbCfg.MaxIdleConns)
		}

		if appDbCfg.MaxOpenConns > 0 {
			db.DB().SetMaxOpenConns(appDbCfg.MaxOpenConns)
		}
		db.LogMode(appDbCfg.Debug)
		return db
	} else {
		log.Fatal(err)
	}
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		pld_logger.Error("db 未初始化")
	}
	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		pld_logger.Error("close db err", err)
	} else {
		log.Println("db closed")
	}
}
