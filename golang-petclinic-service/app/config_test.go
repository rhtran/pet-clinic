package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AppConfigValidation(t *testing.T) {
	config := appConfig{
		AppInfo: AppInfoConfig{
			Name:    "TestApp",
			Version: "1.0.0",
		},
		Database: DatabaseConfig{
			Postgres: PostgresConfig{
				Driver: "postgres",
				Dsn:    "postgres://user:password@localhost:5432/dbname",
			},
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			MaxIdleTime:  30,
		},
		Okta: OktaConfig{
			OAuth2: OAuth2Config{
				URL:      "https://example.com",
				Audience: "audience",
				Issuer:   "issuer",
				ClientId: "clientid",
				Scopes:   "scopes",
			},
		},
		Server: ServerConfig{
			HttpPort: ":8080",
		},
	}

	err := config.Validate()
	assert.NoError(t, err)
}

//func Test_AppConfigValidationMissingAppInfo(t *testing.T) {
//	config := appConfig{
//		Database: DatabaseConfig{
//			Postgres: PostgresConfig{
//				Driver: "postgres",
//				Dsn:    "postgres://user:password@localhost:5432/dbname",
//			},
//			MaxIdleConns: 10,
//			MaxOpenConns: 100,
//			MaxIdleTime:  30,
//		},
//		Okta: OktaConfig{
//			OAuth2: OAuth2Config{
//				URL:      "https://example.com",
//				Audience: "audience",
//				Issuer:   "issuer",
//				ClientId: "clientid",
//				Scopes:   "scopes",
//			},
//		},
//		Server: ServerConfig{
//			HttpPort: ":8080",
//		},
//	}
//
//	err := config.Validate()
//	assert.Error(t, err)
//}
//
//func Test_AppConfigValidationMissingDatabase(t *testing.T) {
//	config := appConfig{
//		AppInfo: AppInfoConfig{
//			Name:    "TestApp",
//			Version: "1.0.0",
//		},
//		Okta: OktaConfig{
//			OAuth2: OAuth2Config{
//				URL:      "https://example.com",
//				Audience: "audience",
//				Issuer:   "issuer",
//				ClientId: "clientid",
//				Scopes:   "scopes",
//			},
//		},
//		Server: ServerConfig{
//			HttpPort: ":8080",
//		},
//	}
//
//	err := config.Validate()
//	assert.Error(t, err)
//}
//
//func Test_AppConfigValidationMissingOkta(t *testing.T) {
//	config := appConfig{
//		AppInfo: AppInfoConfig{
//			Name:    "TestApp",
//			Version: "1.0.0",
//		},
//		Database: DatabaseConfig{
//			Postgres: PostgresConfig{
//				Driver: "postgres",
//				Dsn:    "postgres://user:password@localhost:5432/dbname",
//			},
//			MaxIdleConns: 10,
//			MaxOpenConns: 100,
//			MaxIdleTime:  30,
//		},
//		Server: ServerConfig{
//			HttpPort: ":8080",
//		},
//	}
//
//	err := config.Validate()
//	assert.Error(t, err)
//}
//
//func Test_AppConfigValidationMissingServer(t *testing.T) {
//	config := appConfig{
//		AppInfo: AppInfoConfig{
//			Name:    "TestApp",
//			Version: "1.0.0",
//		},
//		Database: DatabaseConfig{
//			Postgres: PostgresConfig{
//				Driver: "postgres",
//				Dsn:    "postgres://user:password@localhost:5432/dbname",
//			},
//			MaxIdleConns: 10,
//			MaxOpenConns: 100,
//			MaxIdleTime:  30,
//		},
//		Okta: OktaConfig{
//			OAuth2: OAuth2Config{
//				URL:      "https://example.com",
//				Audience: "audience",
//				Issuer:   "issuer",
//				ClientId: "clientid",
//				Scopes:   "scopes",
//			},
//		},
//	}
//
//	err := config.Validate()
//	assert.Error(t, err)
//}
