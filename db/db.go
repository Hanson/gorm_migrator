package db

import (
	"context"
	"github.com/hanson/gorm_migrator"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func InitDb(dsn string) {
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger: logger.New(log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				Colorful:                  true,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				LogLevel:                  logger.Info,
			}),
		NamingStrategy: gorm_migrator.DefaultNaming{},
	})
	if err != nil {
		panic(err)
	}

	m := &gorm_migrator.MyMigrator{}
	m.DB = Db
	m.Migrator.Migrator.Dialector = Db.Dialector
}

func SetDb(db *gorm.DB) {
	Db = db
}

func Ctx(ctx context.Context) *gorm.DB {
	return Db.WithContext(ctx)
}

func NewInstance(m interface{}) *gorm.DB {
	return Db.Model(m).Where("deleted_at = 0")
}
