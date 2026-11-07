package common

import (
	. "aio/config"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Timeline *gorm.DB
}

var Conn DB

func InitDB() {
	Conn.Timeline = OpenDB(Config.Database.Timeline)
}

func OpenDB(database string) *gorm.DB {
	conn, err := gorm.Open(sqlite.Open(database), &gorm.Config{
		//不打印慢日志
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 0,
			},
		),
	})
	if err != nil {
		panic("failed to connect db-")
	}
	sqlDB, _ := conn.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	return conn
}
