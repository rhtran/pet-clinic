package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidDatabaseConfig(t *testing.T) {
	config := DatabaseConfig{
		Postgres: PostgresConfig{
			Driver: "postgres",
			Dsn:    "postgres://user:password@localhost:5432/dbname",
		},
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxIdleTime:  30,
	}

	err := config.Validate()
	assert.NoError(t, err)
}

func Test_InvalidDatabaseConfigMissingPostgres(t *testing.T) {
	config := DatabaseConfig{
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxIdleTime:  30,
	}

	err := config.Validate()
	assert.Error(t, err)
}

func Test_InvalidDatabaseConfigMissingMaxOpenConns(t *testing.T) {
	config := DatabaseConfig{
		Postgres: PostgresConfig{
			Driver: "postgres",
			Dsn:    "postgres://user:password@localhost:5432/dbname",
		},
		MaxIdleConns: 10,
		MaxIdleTime:  30,
	}

	err := config.Validate()
	assert.Error(t, err)
}

func Test_ValidPostgresConfig(t *testing.T) {
	config := PostgresConfig{
		Driver: "postgres",
		Dsn:    "postgres://user:password@localhost:5432/dbname",
	}

	err := config.Validate()
	assert.NoError(t, err)
}

func Test_InvalidPostgresConfigMissingDriver(t *testing.T) {
	config := PostgresConfig{
		Dsn: "postgres://user:password@localhost:5432/dbname",
	}

	err := config.Validate()
	assert.Error(t, err)
}

func Test_InvalidPostgresConfigMissingDsn(t *testing.T) {
	config := PostgresConfig{
		Driver: "postgres",
	}

	err := config.Validate()
	assert.Error(t, err)
}
