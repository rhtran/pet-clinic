package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPetFields(t *testing.T) {
	pet := &Pet{
		Name: "Fido",
		Type: Type{Name: "Dog"},
	}

	assert.Equal(t, "Fido", pet.Name)
	assert.Equal(t, "Dog", pet.Type.Name)
}

func TestPetFieldsWithEmptyName(t *testing.T) {
	pet := &Pet{
		Name: "",
		Type: Type{Name: "Dog"},
	}

	assert.Equal(t, "", pet.Name)
	assert.Equal(t, "Dog", pet.Type.Name)
}

func TestPetFieldsWithEmptyType(t *testing.T) {
	pet := &Pet{
		Name: "Fido",
		Type: Type{Name: ""},
	}

	assert.Equal(t, "Fido", pet.Name)
	assert.Equal(t, "", pet.Type.Name)
}

func (suite *PetRepoTestSuite) TestPetFieldsWithEmptyNameAndType(t *testing.T) {
	pet := &Pet{
		Name: "",
		Type: Type{Name: ""},
	}

	assert.Equal(t, "", pet.Name)
	assert.Equal(t, "", pet.Type.Name)
}
