package pet

import (
	"testing"

	"github.com/qiangxue/go-restful-api/pkg/log"

	"github.com/magiconair/properties/assert"

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
func (petM *petRepositoryMock) FindById(id int) (*Pet, error) {
	args := petM.Called(id)
	intf := args.Get(0)
	val := intf.(Pet)
	ptr := &val

	return ptr, args.Error(1)
}

func (petM *petRepositoryMock) FindByName(name string) ([]Pet, error) {
	args := petM.Called(name)
	intf := args.Get(0)
	val := intf.([]Pet)
	ptr := val

	return ptr, args.Error(1)
}

func (petM *petRepositoryMock) Insert(pet *Pet) (*Pet, error) {
	args := petM.Called(pet)
	intf := args.Get(0)
	val := intf.(*Pet)
	ptr := val

	return ptr, args.Error(1)
}

func (petM *petRepositoryMock) Update(pet *Pet) (*Pet, error) {
	args := petM.Called(pet)
	intf := args.Get(0)
	val := intf.(*Pet)
	ptr := val

	return ptr, args.Error(1)
}

// mocking ends

func Test_GetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetById")
	petMock := petRepositoryMock{}
	pet := &Pet{
		Name: "Nash",
	}
	pet.ID = 1
	petMock.On("FindById", 1).Return(*pet, nil)

	petService := NewPetService(logger, &petMock)
	result, _ := petService.GetPetById(1)
	petMock.AssertExpectations(t)
	petMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, pet.ID, result.ID)
	assert.Equal(t, pet.Name, result.Name)
}

func Test_GetByName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetByName")
	petMock := petRepositoryMock{}
	pet := &Pet{
		Name: "Leo",
	}
	pet.ID = 1

	pets := make([]Pet, 1)
	pets[0] = *pet
	petMock.On("FindByName", "Leo").Return(pets, nil)

	petService := NewPetService(logger, &petMock)
	result, _ := petService.GetPetByName("Leo")
	petMock.AssertExpectations(t)
	petMock.AssertNumberOfCalls(t, "FindByName", 1)

	assert.Equal(t, len(pets), len(result))
	assert.Equal(t, pet.ID, result[0].ID)
	assert.Equal(t, pet.Name, result[0].Name)
}
