// Code generated by mockery v2.42.2. DO NOT EDIT.

package owner

import mock "github.com/stretchr/testify/mock"

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

// create provides a mock function with given fields: owner
func (_m *MockServicer) create(owner *Owner) (*Response, error) {
	ret := _m.Called(owner)

	if len(ret) == 0 {
		panic("no return value specified for create")
	}

	var r0 *Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*Owner) (*Response, error)); ok {
		return rf(owner)
	}
	if rf, ok := ret.Get(0).(func(*Owner) *Response); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*Owner) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'create'
type MockServicer_create_Call struct {
	*mock.Call
}

// create is a helper method to define mock.On call
//   - owner *Owner
func (_e *MockServicer_Expecter) create(owner interface{}) *MockServicer_create_Call {
	return &MockServicer_create_Call{Call: _e.mock.On("create", owner)}
}

func (_c *MockServicer_create_Call) Run(run func(owner *Owner)) *MockServicer_create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Owner))
	})
	return _c
}

func (_c *MockServicer_create_Call) Return(_a0 *Response, _a1 error) *MockServicer_create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_create_Call) RunAndReturn(run func(*Owner) (*Response, error)) *MockServicer_create_Call {
	_c.Call.Return(run)
	return _c
}

// getAllOwners provides a mock function with given fields:
func (_m *MockServicer) getAllOwners() ([]Response, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getAllOwners")
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

// MockServicer_getAllOwners_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getAllOwners'
type MockServicer_getAllOwners_Call struct {
	*mock.Call
}

// getAllOwners is a helper method to define mock.On call
func (_e *MockServicer_Expecter) getAllOwners() *MockServicer_getAllOwners_Call {
	return &MockServicer_getAllOwners_Call{Call: _e.mock.On("getAllOwners")}
}

func (_c *MockServicer_getAllOwners_Call) Run(run func()) *MockServicer_getAllOwners_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockServicer_getAllOwners_Call) Return(_a0 []Response, _a1 error) *MockServicer_getAllOwners_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_getAllOwners_Call) RunAndReturn(run func() ([]Response, error)) *MockServicer_getAllOwners_Call {
	_c.Call.Return(run)
	return _c
}

// getAllOwnersWithPets provides a mock function with given fields:
func (_m *MockServicer) getAllOwnersWithPets() ([]Response, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getAllOwnersWithPets")
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

// MockServicer_getAllOwnersWithPets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getAllOwnersWithPets'
type MockServicer_getAllOwnersWithPets_Call struct {
	*mock.Call
}

// getAllOwnersWithPets is a helper method to define mock.On call
func (_e *MockServicer_Expecter) getAllOwnersWithPets() *MockServicer_getAllOwnersWithPets_Call {
	return &MockServicer_getAllOwnersWithPets_Call{Call: _e.mock.On("getAllOwnersWithPets")}
}

func (_c *MockServicer_getAllOwnersWithPets_Call) Run(run func()) *MockServicer_getAllOwnersWithPets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockServicer_getAllOwnersWithPets_Call) Return(_a0 []Response, _a1 error) *MockServicer_getAllOwnersWithPets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_getAllOwnersWithPets_Call) RunAndReturn(run func() ([]Response, error)) *MockServicer_getAllOwnersWithPets_Call {
	_c.Call.Return(run)
	return _c
}

// getOwnerById provides a mock function with given fields: id
func (_m *MockServicer) getOwnerById(id int) (*Response, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for getOwnerById")
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

// MockServicer_getOwnerById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getOwnerById'
type MockServicer_getOwnerById_Call struct {
	*mock.Call
}

// getOwnerById is a helper method to define mock.On call
//   - id int
func (_e *MockServicer_Expecter) getOwnerById(id interface{}) *MockServicer_getOwnerById_Call {
	return &MockServicer_getOwnerById_Call{Call: _e.mock.On("getOwnerById", id)}
}

func (_c *MockServicer_getOwnerById_Call) Run(run func(id int)) *MockServicer_getOwnerById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockServicer_getOwnerById_Call) Return(_a0 *Response, _a1 error) *MockServicer_getOwnerById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_getOwnerById_Call) RunAndReturn(run func(int) (*Response, error)) *MockServicer_getOwnerById_Call {
	_c.Call.Return(run)
	return _c
}

// getOwnerByLastName provides a mock function with given fields: lastName
func (_m *MockServicer) getOwnerByLastName(lastName string) ([]Response, error) {
	ret := _m.Called(lastName)

	if len(ret) == 0 {
		panic("no return value specified for getOwnerByLastName")
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

// MockServicer_getOwnerByLastName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getOwnerByLastName'
type MockServicer_getOwnerByLastName_Call struct {
	*mock.Call
}

// getOwnerByLastName is a helper method to define mock.On call
//   - lastName string
func (_e *MockServicer_Expecter) getOwnerByLastName(lastName interface{}) *MockServicer_getOwnerByLastName_Call {
	return &MockServicer_getOwnerByLastName_Call{Call: _e.mock.On("getOwnerByLastName", lastName)}
}

func (_c *MockServicer_getOwnerByLastName_Call) Run(run func(lastName string)) *MockServicer_getOwnerByLastName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServicer_getOwnerByLastName_Call) Return(_a0 []Response, _a1 error) *MockServicer_getOwnerByLastName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_getOwnerByLastName_Call) RunAndReturn(run func(string) ([]Response, error)) *MockServicer_getOwnerByLastName_Call {
	_c.Call.Return(run)
	return _c
}

// update provides a mock function with given fields: owner
func (_m *MockServicer) update(owner *Owner) (*Response, error) {
	ret := _m.Called(owner)

	if len(ret) == 0 {
		panic("no return value specified for update")
	}

	var r0 *Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*Owner) (*Response, error)); ok {
		return rf(owner)
	}
	if rf, ok := ret.Get(0).(func(*Owner) *Response); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*Owner) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicer_update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'update'
type MockServicer_update_Call struct {
	*mock.Call
}

// update is a helper method to define mock.On call
//   - owner *Owner
func (_e *MockServicer_Expecter) update(owner interface{}) *MockServicer_update_Call {
	return &MockServicer_update_Call{Call: _e.mock.On("update", owner)}
}

func (_c *MockServicer_update_Call) Run(run func(owner *Owner)) *MockServicer_update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Owner))
	})
	return _c
}

func (_c *MockServicer_update_Call) Return(_a0 *Response, _a1 error) *MockServicer_update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicer_update_Call) RunAndReturn(run func(*Owner) (*Response, error)) *MockServicer_update_Call {
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
