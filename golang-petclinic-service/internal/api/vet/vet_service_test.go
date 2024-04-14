package vet

import (
	vet2 "github.com/rhtran/golang-petclinic-service/pkg/infra/repository/vet"
	model2 "github.com/rhtran/golang-petclinic-service/pkg/model"
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/qiangxue/go-restful-api/pkg/log"
)

func Test_GetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetById")
	vetMock := vet2.MockVetRepositorier{}
	vet := &vet2.Vet{
		Base: model2.Base{
			ID: 1,
		},
		Person: model2.Person{
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
	vetMock := vet2.MockVetRepositorier{}

	vets := make([]vet2.Vet, 1)
	vet := &vet2.Vet{
		Base: model2.Base{
			ID: 1,
		},
		Person: model2.Person{
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
