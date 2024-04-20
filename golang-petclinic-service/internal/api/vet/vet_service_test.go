package vet

import (
	"errors"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"gorm.io/gorm"
	"testing"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

// Test_getById tests the getVetById function
func Test_getById(t *testing.T) {
	testCases := []struct {
		name          string
		expectedVet   *repository.Vet
		expectedError error
	}{
		{
			name: "get vet by id",
			expectedVet: &repository.Vet{
				Person: model.Person{
					FirstName: "Nat",
					LastName:  "Cole",
				},
			},
			expectedError: nil,
		},
		{
			name:          "get vet by id with error",
			expectedVet:   nil,
			expectedError: errors.New("getById: unable to find vet by id"),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_getById")
			vetMock := repository.MockVetRepositorier{}

			vetMock.On("FindById", 1).Return(tc.expectedVet, tc.expectedError)
			vetService := NewVetService(logger, &vetMock)
			result, err := vetService.getVetById(1)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tc.expectedVet.ID, result.ID)
				assert.Equal(t, tc.expectedVet.FirstName, result.FirstName)
				assert.Equal(t, tc.expectedVet.LastName, result.LastName)
				assert.Nil(t, err)
			}

			vetMock.AssertExpectations(t)
			vetMock.AssertNumberOfCalls(t, "FindById", 1)
		})
	}
}

// Test_getByLastName tests the getVetByLastName function
func Test_getByLastName(t *testing.T) {
	testCases := []struct {
		name          string
		expectedVets  []repository.Vet
		expectedError error
	}{
		{
			name: "get vet by last name",
			expectedVets: []repository.Vet{
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
		},
		{
			name:          "get vet by last name with error",
			expectedVets:  nil,
			expectedError: errors.New("getByLastName: unable to find vet by last name"),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_getByLastName")
			vetMock := repository.MockVetRepositorier{}
			vetMock.On("FindByLastName", "DiCaprio").Return(tc.expectedVets, tc.expectedError)

			vetService := NewVetService(logger, &vetMock)
			result, err := vetService.getVetByLastName("DiCaprio")

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				for i, vet := range tc.expectedVets {
					assert.Equal(t, vet.ID, result[i].ID)
					assert.Equal(t, vet.FirstName, result[i].FirstName)
					assert.Equal(t, vet.LastName, result[i].LastName)
					assert.Nil(t, err)
				}
			}
			vetMock.AssertExpectations(t)
			vetMock.AssertNumberOfCalls(t, "FindByLastName", 1)
		})
	}
}
