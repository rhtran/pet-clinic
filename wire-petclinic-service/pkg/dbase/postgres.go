package dbase

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	"wire-petclinic-service/app"
)

type Postgres struct {
}

/*
Create connection pooling using gorm postgres driver.
*/
func PgConnect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(app.Config.Database.Postgres.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.New().Error("Error connecting Postgres database: %v", err)
		panic(err)
	}

	// connection pooling configuration
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(app.Config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(app.Config.Database.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(app.Config.Database.MaxIdleTime) * time.Second)

	return db
}
