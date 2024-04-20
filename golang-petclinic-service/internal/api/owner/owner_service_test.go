package owner

import (
	"errors"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

// Test_getById tests the getOwnerById function
func Test_getById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_getById")
	ownerMock := repository.MockOwnerRepositorier{}
	owner := &repository.Owner{
		Model: gorm.Model{
			ID: 1,
		},
		Person: model.Person{
			FirstName: "Nat",
			LastName:  "Cole",
		},
	}
	ownerMock.On("FindById", 1).Return(owner, nil)

	ownerService := NewOwnerService(logger, &ownerMock)
	result, err := ownerService.getOwnerById(1)
	ownerMock.AssertExpectations(t)
	ownerMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, owner.ID, result.ID)
	assert.Equal(t, owner.FirstName, result.FirstName)
	assert.Equal(t, owner.LastName, result.LastName)
	assert.Nil(t, err)
}

// Test_getById_withError tests the getOwnerById function with an error
func Test_getById_withError(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_getById")
	ownerMock := repository.MockOwnerRepositorier{}
	err := errors.New("getById: unable to find owner by id")
	ownerMock.On("FindById", 1).Return(nil, err)

	ownerService := NewOwnerService(logger, &ownerMock)
	result, err := ownerService.getOwnerById(1)
	ownerMock.AssertExpectations(t)
	ownerMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, "getById: unable to find owner by id", err.Error())
	assert.Nil(t, result)
}

// Test_getByLastName tests the getOwnerByLastName function
func Test_getByLastName(t *testing.T) {
	testCases := []struct {
		name           string
		lastName       string
		expectedOwners []repository.Owner
		expectedError  error
	}{
		{
			name:     "Test get owner by last name success",
			lastName: "DiCaprio",
			expectedOwners: []repository.Owner{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Person: model.Person{
						FirstName: "Leo",
						LastName:  "DiCaprio",
					},
				},
			},
			expectedError: nil,
		},
		{
			name:           "Test get owner by lastname with error",
			lastName:       "Murphy",
			expectedOwners: nil,
			expectedError:  errors.New("getById: unable to find owner by lastName"),
		}, // Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_getByLastName")
			ownerMock := repository.MockOwnerRepositorier{}
			ownerMock.On("FindByLastName", tc.lastName).Return(tc.expectedOwners, tc.expectedError)

			ownerService := NewOwnerService(logger, &ownerMock)
			result, err := ownerService.getOwnerByLastName(tc.lastName)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, len(tc.expectedOwners), len(result))
				for i, owner := range tc.expectedOwners {
					assert.Equal(t, owner.ID, result[i].ID)
					assert.Equal(t, owner.FirstName, result[i].FirstName)
					assert.Equal(t, owner.LastName, result[i].LastName)
					assert.Nil(t, err)
				}
			}

			ownerMock.AssertExpectations(t)
			// Verify that the FindByLastName method is called once
			// for each test case. Therefore, the number of calls = i+1
			ownerMock.AssertNumberOfCalls(t, "FindByLastName", 1)
		})
	}
}

func Test_getAllOwners(t *testing.T) {
	testCases := []struct {
		name           string
		expectedOwners []repository.Owner
		expectedError  error
	}{
		{
			name: "Test get all owners success",
			expectedOwners: []repository.Owner{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Person: model.Person{
						FirstName: "Leo",
						LastName:  "DiCaprio",
					},
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					Person: model.Person{
						FirstName: "Tom",
						LastName:  "Hanks",
					},
				},
			},
			expectedError: nil,
		},
		{
			name:           "Test get all owners with error",
			expectedOwners: nil,
			expectedError:  errors.New("getAllOwners: unable to retrieve owners"),
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_getAllOwners")
			ownerMock := repository.MockOwnerRepositorier{}
			ownerMock.On("FindAll").Return(tc.expectedOwners, tc.expectedError)

			ownerService := NewOwnerService(logger, &ownerMock)
			result, err := ownerService.getAllOwners()

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, len(tc.expectedOwners), len(result))
				for i, owner := range tc.expectedOwners {
					assert.Equal(t, owner.ID, result[i].ID)
					assert.Equal(t, owner.FirstName, result[i].FirstName)
					assert.Equal(t, owner.LastName, result[i].LastName)
					assert.Nil(t, err)
				}
			}

			ownerMock.AssertExpectations(t)
			ownerMock.AssertNumberOfCalls(t, "FindAll", 1)
		})
	}
}

func Test_getAllOwnersWithPets(t *testing.T) {
	testCases := []struct {
		name           string
		expectedOwners []repository.Owner
		expectedError  error
	}{
		{
			name: "Test get all owners with pets success",
			expectedOwners: []repository.Owner{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Person: model.Person{
						FirstName: "Leo",
						LastName:  "DiCaprio",
					},
					Pets: []repository.Pet{
						{
							Name:      "Buddy",
							BirthDate: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
							Type: repository.Type{
								Name: "Dog",
							},
						},
					},
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					Person: model.Person{
						FirstName: "Tom",
						LastName:  "Hanks",
					},
					Pets: []repository.Pet{
						{
							Name:      "Kitty",
							BirthDate: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
							Type: repository.Type{
								Name: "Cat",
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:           "Test get all owners with pets with error",
			expectedOwners: nil,
			expectedError:  errors.New("getAllOwnersWithPets: unable to retrieve owners"),
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_getAllOwnersWithPets")
			ownerMock := repository.MockOwnerRepositorier{}
			ownerMock.On("FindAllWithPets").Return(tc.expectedOwners, tc.expectedError)

			ownerService := NewOwnerService(logger, &ownerMock)
			result, err := ownerService.getAllOwnersWithPets()

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, len(tc.expectedOwners), len(result))
				for i, owner := range tc.expectedOwners {
					assert.Equal(t, owner.ID, result[i].ID)
					assert.Equal(t, owner.FirstName, result[i].FirstName)
					assert.Equal(t, owner.LastName, result[i].LastName)
					assert.Equal(t, len(owner.Pets), len(result[i].Pets))
					assert.Nil(t, err)
				}
			}

			ownerMock.AssertExpectations(t)
			ownerMock.AssertNumberOfCalls(t, "FindAllWithPets", 1)
		})
	}
}

// Test_update tests the update function
func Test_update(t *testing.T) {
	testCases := []struct {
		name          string
		expectedOwner *repository.Owner
		expectedError error
	}{
		{
			name: "Test update owner success",
			expectedOwner: &repository.Owner{
				Model: gorm.Model{
					ID: 1,
				},
				Person: model.Person{
					FirstName: "Nat",
					LastName:  "Cole",
				},
			},
			expectedError: nil,
		},
		{
			name:          "Test update owner with error",
			expectedOwner: nil,
			expectedError: errors.New("update: unable to update owner"),
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_update")
			ownerMock := repository.MockOwnerRepositorier{}
			ownerMock.On("Update", tc.expectedOwner).Return(tc.expectedOwner, tc.expectedError)

			ownerService := NewOwnerService(logger, &ownerMock)
			result, err := ownerService.update(tc.expectedOwner)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tc.expectedOwner.ID, result.ID)
				assert.Equal(t, tc.expectedOwner.FirstName, result.FirstName)
				assert.Equal(t, tc.expectedOwner.LastName, result.LastName)
				assert.Nil(t, err)
			}

			ownerMock.AssertExpectations(t)
			ownerMock.AssertNumberOfCalls(t, "Update", 1)
		})
	}
}

func Test_create(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_create")
	ownerMock := repository.MockOwnerRepositorier{}
	owner := &repository.Owner{
		Model: gorm.Model{
			ID: 1,
		},
		Person: model.Person{
			FirstName: "Nat",
			LastName:  "Cole",
		},
	}
	ownerMock.On("Create", owner).Return(owner, nil)

	ownerService := NewOwnerService(logger, &ownerMock)
	result, err := ownerService.create(owner)
	ownerMock.AssertExpectations(t)
	ownerMock.AssertNumberOfCalls(t, "Create", 1)

	assert.Equal(t, owner.ID, result.ID)
	assert.Equal(t, owner.FirstName, result.FirstName)
	assert.Equal(t, owner.LastName, result.LastName)
	assert.Nil(t, err)
}
