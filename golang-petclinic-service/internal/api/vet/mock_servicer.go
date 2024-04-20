// Code generated by mockery v2.42.2. DO NOT EDIT.

package vet

import (
	repository "github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	mock "github.com/stretchr/testify/mock"
)

// MockServicer is an autogenerated mock type for the Servicer type
type MockServicer struct {
	mock.Mock
}

type MockServicer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServicer) EXPECT() *MockServicer_Expecter {
	return &MockServicer_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: vet
func (_m *MockServicer) create(vet *repository.Vet) (*Response, error) {
	ret := _m.Called(vet)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*repository.Vet) (*Response, error)); ok {
		return rf(vet)
	}
	if rf, ok := ret.Get(0).(func(*repository.Vet) *Response); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*repository.Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockServicer_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - vet *repository.Vet
func (_e *MockServicer_Expecter) create(vet interface{}) *MockServicer_Create_Call {
	return &MockServicer_Create_Call{Call: _e.mock.On("Create", vet)}
}

func (_c *MockServicer_Create_Call) Run(run func(vet *repository.Vet)) *MockServicer_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*repository.Vet))
	})
	return _c
}

func (_c *MockServicer_Create_Call) Return(_a0 *Response, _a1 error) *MockServicer_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_Create_Call) RunAndReturn(run func(*repository.Vet) (*Response, error)) *MockServicer_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllVets provides a mock function with given fields:
func (_m *MockServicer) getAllVets() ([]Response, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllVets")
	}

	var r0 []Response
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Response, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Response)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_GetAllVets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllVets'
type MockServicer_GetAllVets_Call struct {
	*mock.Call
}

// GetAllVets is a helper method to define mock.On call
func (_e *MockServicer_Expecter) getAllVets() *MockServicer_GetAllVets_Call {
	return &MockServicer_GetAllVets_Call{Call: _e.mock.On("GetAllVets")}
}

func (_c *MockServicer_GetAllVets_Call) Run(run func()) *MockServicer_GetAllVets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockServicer_GetAllVets_Call) Return(_a0 []Response, _a1 error) *MockServicer_GetAllVets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_GetAllVets_Call) RunAndReturn(run func() ([]Response, error)) *MockServicer_GetAllVets_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllVetsWithSpecialties provides a mock function with given fields:
func (_m *MockServicer) getAllVetsWithSpecialties() ([]Response, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllVetsWithSpecialties")
	}

	var r0 []Response
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Response, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Response)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_GetAllVetsWithSpecialties_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllVetsWithSpecialties'
type MockServicer_GetAllVetsWithSpecialties_Call struct {
	*mock.Call
}

// GetAllVetsWithSpecialties is a helper method to define mock.On call
func (_e *MockServicer_Expecter) getAllVetsWithSpecialties() *MockServicer_GetAllVetsWithSpecialties_Call {
	return &MockServicer_GetAllVetsWithSpecialties_Call{Call: _e.mock.On("GetAllVetsWithSpecialties")}
}

func (_c *MockServicer_GetAllVetsWithSpecialties_Call) Run(run func()) *MockServicer_GetAllVetsWithSpecialties_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockServicer_GetAllVetsWithSpecialties_Call) Return(_a0 []Response, _a1 error) *MockServicer_GetAllVetsWithSpecialties_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_GetAllVetsWithSpecialties_Call) RunAndReturn(run func() ([]Response, error)) *MockServicer_GetAllVetsWithSpecialties_Call {
	_c.Call.Return(run)
	return _c
}

// GetVetById provides a mock function with given fields: id
func (_m *MockServicer) getVetById(id int) (*Response, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetVetById")
	}

	var r0 *Response
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*Response, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *Response); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_GetVetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVetById'
type MockServicer_GetVetById_Call struct {
	*mock.Call
}

// GetVetById is a helper method to define mock.On call
//   - id int
func (_e *MockServicer_Expecter) getVetById(id interface{}) *MockServicer_GetVetById_Call {
	return &MockServicer_GetVetById_Call{Call: _e.mock.On("GetVetById", id)}
}

func (_c *MockServicer_GetVetById_Call) Run(run func(id int)) *MockServicer_GetVetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockServicer_GetVetById_Call) Return(_a0 *Response, _a1 error) *MockServicer_GetVetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_GetVetById_Call) RunAndReturn(run func(int) (*Response, error)) *MockServicer_GetVetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetVetByLastName provides a mock function with given fields: lastName
func (_m *MockServicer) getVetByLastName(lastName string) ([]Response, error) {
	ret := _m.Called(lastName)

	if len(ret) == 0 {
		panic("no return value specified for GetVetByLastName")
	}

	var r0 []Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]Response, error)); ok {
		return rf(lastName)
	}
	if rf, ok := ret.Get(0).(func(string) []Response); ok {
		r0 = rf(lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_GetVetByLastName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVetByLastName'
type MockServicer_GetVetByLastName_Call struct {
	*mock.Call
}

// GetVetByLastName is a helper method to define mock.On call
//   - lastName string
func (_e *MockServicer_Expecter) getVetByLastName(lastName interface{}) *MockServicer_GetVetByLastName_Call {
	return &MockServicer_GetVetByLastName_Call{Call: _e.mock.On("GetVetByLastName", lastName)}
}

func (_c *MockServicer_GetVetByLastName_Call) Run(run func(lastName string)) *MockServicer_GetVetByLastName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServicer_GetVetByLastName_Call) Return(_a0 []Response, _a1 error) *MockServicer_GetVetByLastName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_GetVetByLastName_Call) RunAndReturn(run func(string) ([]Response, error)) *MockServicer_GetVetByLastName_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: vet
func (_m *MockServicer) update(vet *repository.Vet) (*Response, error) {
	ret := _m.Called(vet)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*repository.Vet) (*Response, error)); ok {
		return rf(vet)
	}
	if rf, ok := ret.Get(0).(func(*repository.Vet) *Response); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*repository.Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockServicer_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - vet *repository.Vet
func (_e *MockServicer_Expecter) update(vet interface{}) *MockServicer_Update_Call {
	return &MockServicer_Update_Call{Call: _e.mock.On("Update", vet)}
}

func (_c *MockServicer_Update_Call) Run(run func(vet *repository.Vet)) *MockServicer_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*repository.Vet))
	})
	return _c
}

func (_c *MockServicer_Update_Call) Return(_a0 *Response, _a1 error) *MockServicer_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_Update_Call) RunAndReturn(run func(*repository.Vet) (*Response, error)) *MockServicer_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServicer creates a new instance of MockServicer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServicer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServicer {
	mock := &MockServicer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
