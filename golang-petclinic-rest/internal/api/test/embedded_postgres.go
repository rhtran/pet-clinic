package test

import (
	"testing"

	"github.com/pressly/goose"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"), &gorm.Config{})
}

func PgStart(t *testing.T, dir string) *embeddedpostgres.EmbeddedPostgres {
	postgresql := embeddedpostgres.NewDatabase(embeddedpostgres.
		DefaultConfig().Port(5433))

	if err := postgresql.Start(); err != nil {
		t.Fatal(err)
	}

	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, _ := db.DB()
	if err := goose.Up(sqlDB, dir); err != nil {
		t.Fatal(err)
	}

	return postgresql
}
