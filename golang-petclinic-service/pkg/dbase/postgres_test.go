package dbase

import (
	"github.com/rhtran/golang-petclinic-service/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgConnectSuccess(t *testing.T) {
	db := PgConnect()

	assert.NotNil(t, db)

	sqlDB, _ := db.DB()
	assert.NotNil(t, sqlDB)

	assert.Equal(t, app.Config.Database.MaxIdleConns, sqlDB.Stats().MaxIdleClosed)
	assert.Equal(t, app.Config.Database.MaxOpenConns, sqlDB.Stats().MaxOpenConnections)
}

//func TestPgConnectPanicOnInvalidDSN(t *testing.T) {
//	// Backup original DSN
//	originalDSN := app.Config.Database.Postgres.Dsn
//
//	// Set invalid DSN
//	app.Config.Database.Postgres.Dsn = "invalid_dsn"
//
//	assert.Panics(t, PgConnect)
//
//	// Restore original DSN
//	app.Config.Database.Postgres.Dsn = originalDSN
//}
//
//func TestPgConnectPanicOnMaxIdleConnsExceeded(t *testing.T) {
//	// Backup original MaxIdleConns
//	originalMaxIdleConns := app.Config.Database.MaxIdleConns
//
//	// Set MaxIdleConns to exceed limit
//	app.Config.Database.MaxIdleConns = sqlDB.Stats().MaxOpenConnections + 1
//
//	assert.Panics(t, PgConnect)
//
//	// Restore original MaxIdleConns
//	app.Config.Database.MaxIdleConns = originalMaxIdleConns
//}
//
//func TestPgConnectPanicOnMaxOpenConnsExceeded(t *testing.T) {
//	// Backup original MaxOpenConns
//	originalMaxOpenConns := app.Config.Database.MaxOpenConns
//
//	// Set MaxOpenConns to exceed limit
//	app.Config.Database.MaxOpenConns = sqlDB.Stats().MaxOpenConnections + 1
//
//	assert.Panics(t, PgConnect)
//
//	// Restore original MaxOpenConns
//	app.Config.Database.MaxOpenConns = originalMaxOpenConns
//}
