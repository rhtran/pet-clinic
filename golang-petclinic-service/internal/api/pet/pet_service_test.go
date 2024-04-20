package pet

import (
	"errors"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"gorm.io/gorm"
	"testing"
	"time"

	"github.com/qiangxue/go-restful-api/pkg/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// for learning purpose, how manual mocking can be done.
// petRepositoryMock
//
//	requires mocking all functions; otherwise, compilation errors.
type petRepositoryMock struct {
	mock.Mock
}

// FindById is a method on the petRepositoryMock struct. It simulates the behavior of
// fetching a Pet object from a data repository based on its unique identifier.
//
// This method takes an integer id as an argument, which is expected to be the unique
// identifier of a Pet object in a data repository.
//
// The method starts by calling the Called method with the id argument. This is a method
// provided by the mock package, which records the fact that a method has been called with
// specific arguments. It returns a mock.Call object.
//
// The mock.Call object has a Get method that retrieves the argument at the specified index.
// In this case, it retrieves the first argument (index 0), which is the Pet object that the
// FindById method is expected to return. This Pet object is then type-asserted to the Pet type.
//
// The method then returns a pointer to the Pet object and an error. The error is expected to be
// nil if the method operates as expected, or an instance of error if something goes wrong.
func (petM *petRepositoryMock) FindById(id int) (*repository.Pet, error) {
	args := petM.Called(id)
	intf := args.Get(0)
	val := intf.(repository.Pet)
	ptr := &val

	return ptr, args.Error(1)
}

func (petM *petRepositoryMock) FindByName(name string) ([]repository.Pet, error) {
	args := petM.Called(name)
	intf := args.Get(0)
	val := intf.([]repository.Pet)
	ptr := val

	return ptr, args.Error(1)
}

func (petM *petRepositoryMock) Insert(pet *repository.Pet) (*repository.Pet, error) {
	args := petM.Called(pet)
	intf := args.Get(0)
	val := intf.(*repository.Pet)
	ptr := val

	return ptr, args.Error(1)
}

func (petM *petRepositoryMock) Update(pet *repository.Pet) (*repository.Pet, error) {
	args := petM.Called(pet)
	intf := args.Get(0)
	val := intf.(*repository.Pet)
	ptr := val

	return ptr, args.Error(1)
}

// mocking ends

func Test_getById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_getById")
	petMock := petRepositoryMock{}
	birthDate := time.Date(2014, 10, 7, 0, 0, 0, 0, time.UTC)
	pet := &repository.Pet{
		Model: gorm.Model{
			ID: 1,
		},
		Name:      "Nash",
		BirthDate: birthDate,
		Type: repository.Type{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Dog",
		},
	}
	petMock.On("FindById", 1).Return(*pet, nil)

	petService := NewPetService(logger, &petMock)
	result, _ := petService.getPetById(1)
	petMock.AssertExpectations(t)
	petMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, pet.ID, result.ID, "Pet ID should be the same")
	assert.Equal(t, pet.Name, result.Name, "Pet Name should be the same")
	assert.Equal(t, pet.BirthDate.Format(time.DateOnly), result.BirthDate, "Pet BirthDate should be the same")
}

func Test_getByName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_getByName")
	petMock := petRepositoryMock{}
	birthDate := time.Date(2017, 7, 2, 0, 0, 0, 0, time.UTC)
	pet := &repository.Pet{
		Model: gorm.Model{
			ID: 1,
		},
		Name:      "Leo",
		BirthDate: birthDate,
		Type: repository.Type{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "Cat",
		},
	}

	pets := make([]repository.Pet, 1)
	pets[0] = *pet
	petMock.On("FindByName", "Leo").Return(pets, nil)

	petService := NewPetService(logger, &petMock)
	result, _ := petService.getPetByName("Leo")
	petMock.AssertExpectations(t)
	petMock.AssertNumberOfCalls(t, "FindByName", 1)

	assert.Equal(t, len(pets), len(result), "Pet length should be the same")
	assert.Equal(t, pet.ID, result[0].ID, "Pet ID should be the same")
	assert.Equal(t, pet.Name, result[0].Name, "Pet Name should be the same")
}

// Test_update is a test function that tests the update method of the PetService struct.
func Test_update(t *testing.T) {
	testCases := []struct {
		name          string
		expectedPet   *repository.Pet
		expectedError error
	}{
		{
			name: "Test update pet success",
			expectedPet: &repository.Pet{
				Model: gorm.Model{
					ID: 1,
				},
				Name:      "Leo",
				BirthDate: time.Date(2017, 7, 2, 0, 0, 0, 0, time.UTC),
				Type: repository.Type{
					Model: gorm.Model{
						ID: 2,
					},
					Name: "Cat",
				},
			},
			expectedError: nil,
		},
		{
			name:          "Test update pet failed",
			expectedPet:   nil,
			expectedError: errors.New("failed to update pet"),
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_update")
			petMock := petRepositoryMock{}
			petMock.On("Update", tc.expectedPet).Return(tc.expectedPet, tc.expectedError)

			petService := NewPetService(logger, &petMock)
			result, err := petService.update(tc.expectedPet)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tc.expectedPet.ID, result.ID, "Pet ID should be the same")
				assert.Equal(t, tc.expectedPet.Name, result.Name, "Pet Name should be the same")
				assert.Nil(t, err)
			}

			petMock.AssertExpectations(t)
			petMock.AssertNumberOfCalls(t, "Update", 1)
		})
	}
}

func Test_create(t *testing.T) {
	testCases := []struct {
		name          string
		expectedPet   *repository.Pet
		expectedError error
	}{
		{
			name: "Test create pet success",
			expectedPet: &repository.Pet{
				Model: gorm.Model{
					ID: 1,
				},
				Name:      "Leo",
				BirthDate: time.Date(2017, 7, 2, 0, 0, 0, 0, time.UTC),
				Type: repository.Type{
					Model: gorm.Model{
						ID: 2,
					},
					Name: "Cat",
				},
			},
			expectedError: nil,
		},
		{
			name:          "Test create pet failed",
			expectedPet:   nil,
			expectedError: errors.New("failed to insert pet"),
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_insert")
			petMock := petRepositoryMock{}
			petMock.On("Insert", tc.expectedPet).Return(tc.expectedPet, tc.expectedError)

			petService := NewPetService(logger, &petMock)
			result, err := petService.create(tc.expectedPet)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tc.expectedPet.ID, result.ID, "Pet ID should be the same")
				assert.Equal(t, tc.expectedPet.Name, result.Name, "Pet Name should be the same")
				assert.Nil(t, err)
			}

			petMock.AssertExpectations(t)
			petMock.AssertNumberOfCalls(t, "Insert", 1)
		})
	}
}
