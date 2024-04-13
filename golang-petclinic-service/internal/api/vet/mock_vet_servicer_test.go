// Code generated by mockery v2.42.2. DO NOT EDIT.

package vet

import mock "github.com/stretchr/testify/mock"

// MockVetServicer is an autogenerated mock type for the VetServicer type
type MockVetServicer struct {
	mock.Mock
}

type MockVetServicer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVetServicer) EXPECT() *MockVetServicer_Expecter {
	return &MockVetServicer_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: vet
func (_m *MockVetServicer) Create(vet *Vet) (*Response, error) {
	ret := _m.Called(vet)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*Vet) (*Response, error)); ok {
		return rf(vet)
	}
	if rf, ok := ret.Get(0).(func(*Vet) *Response); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetServicer_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockVetServicer_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - vet *Vet
func (_e *MockVetServicer_Expecter) Create(vet interface{}) *MockVetServicer_Create_Call {
	return &MockVetServicer_Create_Call{Call: _e.mock.On("Create", vet)}
}

func (_c *MockVetServicer_Create_Call) Run(run func(vet *Vet)) *MockVetServicer_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Vet))
	})
	return _c
}

func (_c *MockVetServicer_Create_Call) Return(_a0 *Response, _a1 error) *MockVetServicer_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetServicer_Create_Call) RunAndReturn(run func(*Vet) (*Response, error)) *MockVetServicer_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllVets provides a mock function with given fields:
func (_m *MockVetServicer) GetAllVets() ([]Response, error) {
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

// MockVetServicer_GetAllVets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllVets'
type MockVetServicer_GetAllVets_Call struct {
	*mock.Call
}

// GetAllVets is a helper method to define mock.On call
func (_e *MockVetServicer_Expecter) GetAllVets() *MockVetServicer_GetAllVets_Call {
	return &MockVetServicer_GetAllVets_Call{Call: _e.mock.On("GetAllVets")}
}

func (_c *MockVetServicer_GetAllVets_Call) Run(run func()) *MockVetServicer_GetAllVets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVetServicer_GetAllVets_Call) Return(_a0 []Response, _a1 error) *MockVetServicer_GetAllVets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetServicer_GetAllVets_Call) RunAndReturn(run func() ([]Response, error)) *MockVetServicer_GetAllVets_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllVetsWithSpecialties provides a mock function with given fields:
func (_m *MockVetServicer) GetAllVetsWithSpecialties() ([]Response, error) {
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

// MockVetServicer_GetAllVetsWithSpecialties_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllVetsWithSpecialties'
type MockVetServicer_GetAllVetsWithSpecialties_Call struct {
	*mock.Call
}

// GetAllVetsWithSpecialties is a helper method to define mock.On call
func (_e *MockVetServicer_Expecter) GetAllVetsWithSpecialties() *MockVetServicer_GetAllVetsWithSpecialties_Call {
	return &MockVetServicer_GetAllVetsWithSpecialties_Call{Call: _e.mock.On("GetAllVetsWithSpecialties")}
}

func (_c *MockVetServicer_GetAllVetsWithSpecialties_Call) Run(run func()) *MockVetServicer_GetAllVetsWithSpecialties_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVetServicer_GetAllVetsWithSpecialties_Call) Return(_a0 []Response, _a1 error) *MockVetServicer_GetAllVetsWithSpecialties_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetServicer_GetAllVetsWithSpecialties_Call) RunAndReturn(run func() ([]Response, error)) *MockVetServicer_GetAllVetsWithSpecialties_Call {
	_c.Call.Return(run)
	return _c
}

// GetVetById provides a mock function with given fields: id
func (_m *MockVetServicer) GetVetById(id int) (*Response, error) {
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

// MockVetServicer_GetVetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVetById'
type MockVetServicer_GetVetById_Call struct {
	*mock.Call
}

// GetVetById is a helper method to define mock.On call
//   - id int
func (_e *MockVetServicer_Expecter) GetVetById(id interface{}) *MockVetServicer_GetVetById_Call {
	return &MockVetServicer_GetVetById_Call{Call: _e.mock.On("GetVetById", id)}
}

func (_c *MockVetServicer_GetVetById_Call) Run(run func(id int)) *MockVetServicer_GetVetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockVetServicer_GetVetById_Call) Return(_a0 *Response, _a1 error) *MockVetServicer_GetVetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetServicer_GetVetById_Call) RunAndReturn(run func(int) (*Response, error)) *MockVetServicer_GetVetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetVetByLastName provides a mock function with given fields: lastName
func (_m *MockVetServicer) GetVetByLastName(lastName string) ([]Response, error) {
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

// MockVetServicer_GetVetByLastName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVetByLastName'
type MockVetServicer_GetVetByLastName_Call struct {
	*mock.Call
}

// GetVetByLastName is a helper method to define mock.On call
//   - lastName string
func (_e *MockVetServicer_Expecter) GetVetByLastName(lastName interface{}) *MockVetServicer_GetVetByLastName_Call {
	return &MockVetServicer_GetVetByLastName_Call{Call: _e.mock.On("GetVetByLastName", lastName)}
}

func (_c *MockVetServicer_GetVetByLastName_Call) Run(run func(lastName string)) *MockVetServicer_GetVetByLastName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVetServicer_GetVetByLastName_Call) Return(_a0 []Response, _a1 error) *MockVetServicer_GetVetByLastName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetServicer_GetVetByLastName_Call) RunAndReturn(run func(string) ([]Response, error)) *MockVetServicer_GetVetByLastName_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: vet
func (_m *MockVetServicer) Update(vet *Vet) (*Response, error) {
	ret := _m.Called(vet)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*Vet) (*Response, error)); ok {
		return rf(vet)
	}
	if rf, ok := ret.Get(0).(func(*Vet) *Response); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetServicer_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockVetServicer_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - vet *Vet
func (_e *MockVetServicer_Expecter) Update(vet interface{}) *MockVetServicer_Update_Call {
	return &MockVetServicer_Update_Call{Call: _e.mock.On("Update", vet)}
}

func (_c *MockVetServicer_Update_Call) Run(run func(vet *Vet)) *MockVetServicer_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Vet))
	})
	return _c
}

func (_c *MockVetServicer_Update_Call) Return(_a0 *Response, _a1 error) *MockVetServicer_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetServicer_Update_Call) RunAndReturn(run func(*Vet) (*Response, error)) *MockVetServicer_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVetServicer creates a new instance of MockVetServicer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVetServicer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVetServicer {
	mock := &MockVetServicer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}