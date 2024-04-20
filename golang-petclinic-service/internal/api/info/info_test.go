package info

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo_Validate(t *testing.T) {
	info := Info{
		AppName:     "api",
		Description: "api description",
		Version:     "1.0.0",
	}

	assert.Nil(t, info.Validate())
}

func TestInfo_Invalidate(t *testing.T) {
	infoTable := []Info{{AppName: "api", Description: "api description"}, {Description: "api description"}, {Version: "1.0.0"}}

	for i := range infoTable {
		assert.Error(t, infoTable[i].Validate())
	}
}

func Test_InfoValidationFailsMissingAppName(t *testing.T) {
	info := Info{
		Description: "api description",
		Version:     "1.0.0",
	}

	assert.NotNil(t, info.Validate())
}

func Test_InfoValidationFailsMissingDescription(t *testing.T) {
	info := Info{
		AppName: "api",
		Version: "1.0.0",
	}

	assert.NotNil(t, info.Validate())
}

func Test_InfoValidationFailsMissingVersion(t *testing.T) {
	info := Info{
		AppName:     "api",
		Description: "api description",
	}

	assert.NotNil(t, info.Validate())
}
