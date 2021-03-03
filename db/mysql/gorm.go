package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

// dbDns data sources
// MaxIdle 最大空闲连接数
// MaxOpen 最大打开连接数
// MaxLifeTime 最长活跃时间
func Connect(dbDns string, MaxIdle, MaxOpen, MaxLifeTime int) error {
	sqlDB, err := sql.Open("mysql", dbDns)
	if err != nil {
		return err
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(MaxIdle)
	sqlDB.SetMaxOpenConns(MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(MaxLifeTime))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 启用彩色打印
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})

	if err != nil {
		return err
	}
	fmt.Println("数据库链接成功")
	db.Logger.LogMode(logger.Info)
	return nil
}

func Client() *gorm.DB {
	return db
}
