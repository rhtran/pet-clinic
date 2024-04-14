// Code generated by mockery v2.42.2. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockVisitRepositorier is an autogenerated mock type for the VisitRepositorier type
type MockVisitRepositorier struct {
	mock.Mock
}

type MockVisitRepositorier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVisitRepositorier) EXPECT() *MockVisitRepositorier_Expecter {
	return &MockVisitRepositorier_Expecter{mock: &_m.Mock}
}

// FindAll provides a mock function with given fields:
func (_m *MockVisitRepositorier) FindAll() ([]Visit, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []Visit
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Visit, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Visit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Visit)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVisitRepositorier_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockVisitRepositorier_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockVisitRepositorier_Expecter) FindAll() *MockVisitRepositorier_FindAll_Call {
	return &MockVisitRepositorier_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockVisitRepositorier_FindAll_Call) Run(run func()) *MockVisitRepositorier_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVisitRepositorier_FindAll_Call) Return(_a0 []Visit, _a1 error) *MockVisitRepositorier_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVisitRepositorier_FindAll_Call) RunAndReturn(run func() ([]Visit, error)) *MockVisitRepositorier_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *MockVisitRepositorier) FindById(id int) (*Visit, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *Visit
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*Visit, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *Visit); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Visit)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVisitRepositorier_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockVisitRepositorier_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id int
func (_e *MockVisitRepositorier_Expecter) FindById(id interface{}) *MockVisitRepositorier_FindById_Call {
	return &MockVisitRepositorier_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *MockVisitRepositorier_FindById_Call) Run(run func(id int)) *MockVisitRepositorier_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockVisitRepositorier_FindById_Call) Return(_a0 *Visit, _a1 error) *MockVisitRepositorier_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVisitRepositorier_FindById_Call) RunAndReturn(run func(int) (*Visit, error)) *MockVisitRepositorier_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVisitRepositorier creates a new instance of MockVisitRepositorier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVisitRepositorier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVisitRepositorier {
	mock := &MockVisitRepositorier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
