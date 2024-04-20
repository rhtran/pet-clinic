package visit

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"gorm.io/gorm"
	"testing"
	"time"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

func Test_GetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetById")
	visitMock := repository.MockVisitRepositorier{}
	birthDate := time.Date(2017, 7, 2, 0, 0, 0, 0, time.UTC)
	visit := &repository.Visit{
		Model: gorm.Model{
			ID: 1,
		},
		PetID:       7,
		Pet:         repository.Pet{Model: gorm.Model{ID: 7}, Name: "Leo", BirthDate: &birthDate, Type: repository.Type{Model: gorm.Model{ID: 2}, Name: "Cat"}},
		VisitDate:   time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC),
		Description: "rabies shot",
	}
	visitMock.On("FindById", 1).Return(visit, nil)

	visitService := NewVisitService(logger, &visitMock)
	result, _ := visitService.getVisitById(1)
	visitMock.AssertExpectations(t)
	visitMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, visit.ID, result.ID)
	assert.Equal(t, visit.Description, result.Description)
	assert.Equal(t, visit.Pet.ID, result.PetResponse.ID)
	assert.Equal(t, visit.Pet.Name, result.PetResponse.Name)
	assert.Equal(t, visit.Pet.Type.Name, result.PetResponse.Type)
	assert.Equal(t, visit.Pet.BirthDate.Format(time.DateOnly), result.PetResponse.BirthDate)
}
