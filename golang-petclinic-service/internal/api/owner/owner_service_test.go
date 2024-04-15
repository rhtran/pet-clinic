package owner

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_getById")
	ownerMock := repository.MockOwnerRepositorier{}
	owner := &repository.Owner{
		Base: model.Base{
			ID: 1,
		},
		Person: model.Person{
			FirstName: "Nat",
			LastName:  "Cole",
		},
	}
	ownerMock.On("FindById", 1).Return(owner, nil)

	ownerService := NewOwnerService(logger, &ownerMock)
	result, _ := ownerService.getOwnerById(1)
	ownerMock.AssertExpectations(t)
	ownerMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, owner.ID, result.ID)
	assert.Equal(t, owner.FirstName, result.FirstName)
	assert.Equal(t, owner.LastName, result.LastName)
}

func Test_getByLastName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_getByLastName")
	ownerMock := repository.MockOwnerRepositorier{}

	owners := make([]repository.Owner, 1)
	owner := &repository.Owner{
		Base: model.Base{
			ID: 1,
		},
		Person: model.Person{
			FirstName: "Leo",
			LastName:  "DiCaprio",
		},
	}
	owner.ID = 1
	owners[0] = *owner

	ownerMock.On("FindByLastName", "DiCaprio").Return(owners, nil)
	ownerService := NewOwnerService(logger, &ownerMock)
	result, _ := ownerService.getOwnerByLastName("DiCaprio")
	ownerMock.AssertExpectations(t)
	ownerMock.AssertNumberOfCalls(t, "FindByLastName", 1)

	assert.Equal(t, owner.ID, result[0].ID)
	assert.Equal(t, owner.FirstName, result[0].FirstName)
	assert.Equal(t, owner.LastName, result[0].LastName)
}
