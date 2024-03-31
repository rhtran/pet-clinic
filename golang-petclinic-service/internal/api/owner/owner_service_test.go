package owner

import (
	model2 "github.com/rhtran/golang-petclinic-service/pkg/model"
	"testing"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

func Test_GetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetById")
	ownerMock := MockOwnerRepositorier{}
	owner := &Owner{
		Base: model2.Base{
			ID: 1,
		},
		Person: model2.Person{
			FirstName: "Nat",
			LastName:  "Cole",
		},
	}
	ownerMock.On("FindById", 1).Return(owner, nil)

	ownerService := NewOwnerService(logger, &ownerMock)
	result, _ := ownerService.GetOwnerById(1)
	ownerMock.AssertExpectations(t)
	ownerMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, owner.ID, result.ID)
	assert.Equal(t, owner.FirstName, result.FirstName)
	assert.Equal(t, owner.LastName, result.LastName)
}

func Test_GetByLastName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_GetByLastName")
	ownerMock := MockOwnerRepositorier{}

	owners := make([]Owner, 1)
	owner := &Owner{
		Base: model2.Base{
			ID: 1,
		},
		Person: model2.Person{
			FirstName: "Leo",
			LastName:  "DiCaprio",
		},
	}
	owner.ID = 1
	owners[0] = *owner

	ownerMock.On("FindByLastName", "DiCaprio").Return(owners, nil)
	ownerService := NewOwnerService(logger, &ownerMock)
	result, _ := ownerService.GetOwnerByLastName("DiCaprio")
	ownerMock.AssertExpectations(t)
	ownerMock.AssertNumberOfCalls(t, "FindByLastName", 1)

	assert.Equal(t, owner.ID, result[0].ID)
	assert.Equal(t, owner.FirstName, result[0].FirstName)
	assert.Equal(t, owner.LastName, result[0].LastName)
}
