package mysql

import (
	"core/config"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func New(conf *config.DB, opts ...Option) (DB *gorm.DB, err error) {
	params := NewOptions(opts...)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: params.SlowThreshold, // 慢 SQL 阈值
			LogLevel:      params.LogLevel,      // Log level
			Colorful:      params.Colorful,      // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Prefix,          //设置表前缀
			SingularTable: params.SingularTable, //单数模式 不加s
		},
		Logger: newLogger,
	})
	if err != nil {
		return
	}
	err = db.Use(&GormTracing{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(params.SetMaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(params.SetMaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(params.SetConnMaxLifetime)

	return db, nil
}

// WithTablePrefix 配置表前缀
func WithTablePrefix(prefix string) Option {
	return func(options *Options) {
		options.TablePrefix = prefix
	}
}
