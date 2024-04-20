package visit

import (
	"errors"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_getById(t *testing.T) {
	testCases := []struct {
		name          string
		expectedVisit *repository.Visit
		expectedError error
	}{
		{
			name: "Test get visit by id success",
			expectedVisit: &repository.Visit{
				Model: gorm.Model{
					ID: 1,
				},
				PetID: 7,
				Pet: repository.Pet{
					Model:     gorm.Model{ID: 7},
					Name:      "Leo",
					BirthDate: time.Date(2017, 7, 2, 0, 0, 0, 0, time.UTC),
					Type:      repository.Type{Model: gorm.Model{ID: 2}, Name: "Cat"}},
				VisitDate:   time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC),
				Description: "rabies shot",
			},
			expectedError: nil,
		},
		{
			name:          "Test get visit by id with error",
			expectedVisit: nil,
			expectedError: errors.New("getById: unable to find visit by id"),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_getById")
			visitMock := repository.MockVisitRepositorier{}

			visitMock.On("FindById", 1).Return(tc.expectedVisit, tc.expectedError)
			visitService := NewVisitService(logger, &visitMock)
			result, err := visitService.getVisitById(1)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tc.expectedVisit.ID, result.ID)
				assert.Equal(t, tc.expectedVisit.Description, result.Description)
				assert.Equal(t, tc.expectedVisit.Pet.ID, result.PetResponse.ID)
				assert.Equal(t, tc.expectedVisit.Pet.Name, result.PetResponse.Name)
				assert.Equal(t, tc.expectedVisit.Pet.Type.Name, result.PetResponse.Type)
				assert.Equal(t, tc.expectedVisit.VisitDate.Format(time.DateOnly), result.VisitDate)
				assert.Equal(t, tc.expectedVisit.Pet.BirthDate.Format(time.DateOnly), result.PetResponse.BirthDate)
				assert.Nil(t, err)
			}

			visitMock.AssertExpectations(t)
			visitMock.AssertNumberOfCalls(t, "FindById", 1)
		})
	}
}

// Test_getAllVisits tests the getAllVisits function
func Test_getAllVisits(t *testing.T) {
	testCases := []struct {
		name           string
		expectedVisits []repository.Visit
		expectedError  error
	}{
		{
			name: "Test get all visits success",
			expectedVisits: []repository.Visit{
				{
					Model: gorm.Model{
						ID: 1,
					},
					PetID: 7,
					Pet: repository.Pet{
						Model:     gorm.Model{ID: 7},
						Name:      "Leo",
						BirthDate: time.Date(2017, 7, 2, 0, 0, 0, 0, time.UTC),
						Type:      repository.Type{Model: gorm.Model{ID: 2}, Name: "Cat"}},
					VisitDate:   time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC),
					Description: "rabies shot",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					PetID: 8,
					Pet: repository.Pet{
						Model:     gorm.Model{ID: 8},
						Name:      "Jack",
						BirthDate: time.Date(2018, 8, 2, 0, 0, 0, 0, time.UTC),
						Type:      repository.Type{Model: gorm.Model{ID: 3}, Name: "Dog"}},
					VisitDate:   time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC),
					Description: "rabies shot",
				},
			},
			expectedError: nil,
		},
		{
			name:           "Test get all visits with error",
			expectedVisits: nil,
			expectedError:  errors.New("getAllVisits: unable to find all visits"),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New().With(nil, "function", "Test_getAllVisits")
			visitMock := repository.MockVisitRepositorier{}

			visitMock.On("FindAll").Return(tc.expectedVisits, tc.expectedError)
			visitService := NewVisitService(logger, &visitMock)
			result, err := visitService.getAllVisits()

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, result)
			} else {
				for i, visit := range tc.expectedVisits {
					assert.Equal(t, visit.ID, result[i].ID)
					assert.Equal(t, visit.Description, result[i].Description)
					assert.Equal(t, visit.Pet.ID, result[i].PetResponse.ID)
					assert.Equal(t, visit.Pet.Name, result[i].PetResponse.Name)
					assert.Equal(t, visit.Pet.Type.Name, result[i].PetResponse.Type)
					assert.Equal(t, visit.VisitDate.Format(time.DateOnly), result[i].VisitDate)
					assert.Equal(t, visit.Pet.BirthDate.Format(time.DateOnly), result[i].PetResponse.BirthDate)
					assert.Nil(t, err)
				}
			}

			visitMock.AssertExpectations(t)
			visitMock.AssertNumberOfCalls(t, "FindAll", 1)
		})
	}
}
