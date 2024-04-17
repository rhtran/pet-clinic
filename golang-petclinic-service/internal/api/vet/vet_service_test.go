package vet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/qiangxue/go-restful-api/pkg/log"
)

func Test_GetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetById")
	vetMock := repository.MockVetRepositorier{}
	vet := &repository.Vet{
		Person: model.Person{
			FirstName: "Nat",
			LastName:  "Cole",
		},
	}
	vetMock.On("FindById", 1).Return(vet, nil)

	vetService := NewVetService(logger, &vetMock)
	result, _ := vetService.GetVetById(1)
	vetMock.AssertExpectations(t)
	vetMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, vet.ID, result.ID)
	assert.Equal(t, vet.FirstName, result.FirstName)
	assert.Equal(t, vet.LastName, result.LastName)
}

func Test_GetByLastName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetByLastName")
	vetMock := repository.MockVetRepositorier{}

	vets := make([]repository.Vet, 1)
	vet := &repository.Vet{
		Person: model.Person{
			FirstName: "Leo",
			LastName:  "DiCaprio",
		},
	}
	vets[0] = *vet

	vetMock.On("FindByLastName", "DiCaprio").Return(vets, nil)
	vetService := NewVetService(logger, &vetMock)
	result, _ := vetService.GetVetByLastName("DiCaprio")
	vetMock.AssertExpectations(t)
	vetMock.AssertNumberOfCalls(t, "FindByLastName", 1)

	assert.Equal(t, vet.ID, result[0].ID)
	assert.Equal(t, vet.FirstName, result[0].FirstName)
	assert.Equal(t, vet.LastName, result[0].LastName)
}
