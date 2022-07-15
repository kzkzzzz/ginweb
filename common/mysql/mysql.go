package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type Config struct {
	Dsn         string
	MaxConn     int
	MaxIdleConn int
	MaxLifetime int
	Debug       bool
}

func NewDB(conf *Config) (db *gorm.DB) {
	var err error
	var level = gormLogger.Warn
	if conf.Debug {
		level = gormLogger.Info
	}

	db, err = gorm.Open(mysql.Open(conf.Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 gormLogger.Default.LogMode(level),
	})

	if err != nil {
		panic(fmt.Sprintf("连接mysql失败(-1): %s", err))
	}
	s, _ := db.DB()
	err = s.Ping()
	if err != nil {
		panic(fmt.Sprintf("连接mysql失败(-2): %s", err))
	}

	s.SetMaxOpenConns(conf.MaxConn)
	s.SetMaxIdleConns(conf.MaxIdleConn)
	s.SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Second)
	return
}

func NewGormLogger(writer gormLogger.Writer, level gormLogger.LogLevel) gormLogger.Interface {
	return gormLogger.New(writer, gormLogger.Config{
		SlowThreshold:             time.Second * 2, // 慢查询阈值
		Colorful:                  true,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  level,
	})
}
