package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckStatusIsUp(t *testing.T) {
	hc := Check{Status: "UP"}
	err := hc.Validate()
	assert.Nil(t, err)
}

func Test_CheckStatusIsDown(t *testing.T) {
	hc := Check{Status: "DOWN"}
	err := hc.Validate()
	assert.Nil(t, err)
}

func Test_CheckStatusIsEmpty(t *testing.T) {
	hc := Check{Status: ""}
	err := hc.Validate()
	assert.NotNil(t, err)
}

func Test_CheckStatusIsInvalid(t *testing.T) {
	hc := Check{Status: "INVALID"}
	err := hc.Validate()
	assert.Nil(t, err)
}
