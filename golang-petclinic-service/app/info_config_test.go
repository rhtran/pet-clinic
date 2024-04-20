package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppInfoConfigValidation(t *testing.T) {
	t.Run("ValidAppInfoConfig", func(t *testing.T) {
		config := AppInfoConfig{
			Name:        "TestApp",
			Description: "This is a test app",
			Version:     "1.0.0",
		}

		err := config.Validate()
		assert.NoError(t, err)
	})

	t.Run("InvalidAppInfoConfigMissingName", func(t *testing.T) {
		config := AppInfoConfig{
			Description: "This is a test app",
			Version:     "1.0.0",
		}

		err := config.Validate()
		assert.Error(t, err)
	})

	t.Run("InvalidAppInfoConfigMissingDescription", func(t *testing.T) {
		config := AppInfoConfig{
			Name:    "TestApp",
			Version: "1.0.0",
		}

		err := config.Validate()
		assert.Error(t, err)
	})

	t.Run("InvalidAppInfoConfigMissingVersion", func(t *testing.T) {
		config := AppInfoConfig{
			Name:        "TestApp",
			Description: "This is a test app",
		}

		err := config.Validate()
		assert.Error(t, err)
	})
}
