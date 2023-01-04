package mysql

import (
	"time"

	"gorm.io/gorm/logger"
)

type Option func(*Options)

type Options struct {
	SlowThreshold      time.Duration   // 慢 SQL 阈值
	LogLevel           logger.LogLevel // Log level
	Colorful           bool            // 禁用彩色打印
	TablePrefix        string          //设置表前缀
	SingularTable      bool            //单数模式 不加s
	SetMaxIdleConns    int             //设置空闲连接池中连接的最大数量
	SetMaxOpenConns    int             //设置空闲连接池中连接的最大数量
	SetConnMaxLifetime time.Duration   //设置了连接可复用的最大时间。
}

func defaultOptions() *Options {
	return &Options{
		SlowThreshold:      time.Second,
		LogLevel:           logger.Info,
		Colorful:           true,
		TablePrefix:        "",
		SingularTable:      true,
		SetMaxIdleConns:    25,
		SetMaxOpenConns:    25,
		SetConnMaxLifetime: time.Hour,
	}
}

func NewOptions(opts ...Option) *Options {
	options := defaultOptions()
	for _, o := range opts {
		o(options)
	}
	return options
}
