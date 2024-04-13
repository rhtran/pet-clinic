package visit

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"testing"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

func Test_GetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetById")
	visitMock := MockRepositorier{}
	visit := &Visit{
		Base: model.Base{
			ID: 1,
		},
		PetID:       7,
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
	assert.Equal(t, visit.Pet.BirthDate, result.PetResponse.BirthDate)
}
