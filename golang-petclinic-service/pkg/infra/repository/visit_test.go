package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVisitFields(t *testing.T) {
	visit := &Visit{
		VisitDate:   time.Now(),
		Description: "Regular checkup",
		PetID:       1,
	}

	assert.Equal(t, "Regular checkup", visit.Description)
	assert.Equal(t, 1, visit.PetID)
}

func TestVisitFieldsWithEmptyDescription(t *testing.T) {
	visit := &Visit{
		VisitDate:   time.Now(),
		Description: "",
		PetID:       1,
	}

	assert.Equal(t, "", visit.Description)
	assert.Equal(t, 1, visit.PetID)
}

func TestVisitFieldsWithNoPetID(t *testing.T) {
	visit := &Visit{
		VisitDate:   time.Now(),
		Description: "Regular checkup",
		PetID:       0,
	}

	assert.Equal(t, "Regular checkup", visit.Description)
	assert.Equal(t, 0, visit.PetID)
}

func TestVisitFieldsWithEmptyDescriptionAndNoPetID(t *testing.T) {
	visit := &Visit{
		VisitDate:   time.Now(),
		Description: "",
		PetID:       0,
	}

	assert.Equal(t, "", visit.Description)
	assert.Equal(t, 0, visit.PetID)
}
