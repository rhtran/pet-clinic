package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_Validate(t *testing.T) {
	healthCheck := Check{
		Status: "UP",
	}

	assert.Nil(t, healthCheck.Validate())
}

func TestHeathCheck_Invalidate(t *testing.T) {
	healthCheckTable := []Check{{}}

	for i := range healthCheckTable {
		assert.Error(t, healthCheckTable[i].Validate())
	}
}
