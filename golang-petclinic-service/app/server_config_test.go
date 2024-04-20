package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerConfigValidationWithValidHttpPort(t *testing.T) {
	config := ServerConfig{
		HttpPort: ":8080",
	}

	err := config.Validate()
	assert.NoError(t, err)
}

func TestServerConfigValidationWithEmptyHttpPort(t *testing.T) {
	config := ServerConfig{}

	err := config.Validate()
	assert.Error(t, err)
	assert.EqualError(t, err, "HttpPort: cannot be blank.")
}

func TestServerConfigValidationWithInvalidHttpPort(t *testing.T) {
	t.Run("InvalidPortMissingColon", func(t *testing.T) {
		config := ServerConfig{
			HttpPort: "8080",
		}

		err := config.Validate()
		assert.Error(t, err)
		assert.EqualError(t, err, "HttpPort: must be in the format :dddd.")
	})

	t.Run("InvalidPortMissingColonAndChars", func(t *testing.T) {
		config := ServerConfig{
			HttpPort: "invalid",
		}

		err := config.Validate()
		assert.Error(t, err)
	})

	t.Run("InvalidPortWithChars", func(t *testing.T) {
		config := ServerConfig{
			HttpPort: ":invalid",
		}

		err := config.Validate()
		assert.Error(t, err)
		assert.EqualError(t, err, "HttpPort: must be in the format :dddd.")
	})
}
