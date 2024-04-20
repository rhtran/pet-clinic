package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidOktaConfig(t *testing.T) {
	config := OktaConfig{
		OAuth2: OAuth2Config{
			URL:      "https://example.com",
			Audience: "audience",
			Issuer:   "issuer",
			ClientId: "clientid",
			Scopes:   "scopes",
		},
	}

	err := config.OAuth2.Validate()
	assert.NoError(t, err)
}

func Test_InvalidOktaConfigMissingURL(t *testing.T) {
	config := OktaConfig{
		OAuth2: OAuth2Config{
			Audience: "audience",
			Issuer:   "issuer",
			ClientId: "clientid",
			Scopes:   "scopes",
		},
	}

	err := config.OAuth2.Validate()
	assert.Error(t, err)
}

func Test_InvalidOktaConfigMissingIssuer(t *testing.T) {
	config := OktaConfig{
		OAuth2: OAuth2Config{
			URL:      "https://example.com",
			Audience: "audience",
			ClientId: "clientid",
			Scopes:   "scopes",
		},
	}

	err := config.OAuth2.Validate()
	assert.Error(t, err)
}
