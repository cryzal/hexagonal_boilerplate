package mysql

import (
	"hexagonal_boilerplate/shared/config"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host        string
	User        string
	Password    string
	DBName      string
	DBNumber    int
	Port        int
	DebugMode   bool
	MaxConn     int
	MaxIdle     int
	MaxLifetime int
}

func Connect(cfg *config.Config) *gorm.DB {
	var connection *gorm.DB
	var err error

	config := DBConfig{
		Host:        cfg.Database.Mysql.Host,
		User:        cfg.Database.Mysql.Username,
		Password:    cfg.Database.Mysql.Password,
		DBName:      cfg.Database.Mysql.DBName,
		Port:        cfg.Database.Mysql.Port,
		DebugMode:   cfg.Database.Mysql.Debug,
		MaxConn:     cfg.Database.Mysql.MaxOpenConn,
		MaxIdle:     cfg.Database.Mysql.MaxIdleConn,
		MaxLifetime: cfg.Database.Mysql.MaxLifetime,
	}

	var isDebug = logger.Error

	if config.DebugMode {
		isDebug = logger.Info
	}

	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DBName + "?charset=utf8&parseTime=True&loc=Local"
	connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(isDebug),
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := connection.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(config.MaxConn)
	sqlDB.SetMaxIdleConns(config.MaxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)

	return connection
}
