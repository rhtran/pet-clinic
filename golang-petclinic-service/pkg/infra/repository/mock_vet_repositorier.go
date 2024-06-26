// Code generated by mockery v2.42.2. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockVetRepositorier is an autogenerated mock type for the VetRepositorier type
type MockVetRepositorier struct {
	mock.Mock
}

type MockVetRepositorier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVetRepositorier) EXPECT() *MockVetRepositorier_Expecter {
	return &MockVetRepositorier_Expecter{mock: &_m.Mock}
}

// FindAll provides a mock function with given fields:
func (_m *MockVetRepositorier) FindAll() ([]Vet, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []Vet
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Vet, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Vet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Vet)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetRepositorier_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockVetRepositorier_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockVetRepositorier_Expecter) FindAll() *MockVetRepositorier_FindAll_Call {
	return &MockVetRepositorier_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockVetRepositorier_FindAll_Call) Run(run func()) *MockVetRepositorier_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVetRepositorier_FindAll_Call) Return(_a0 []Vet, _a1 error) *MockVetRepositorier_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetRepositorier_FindAll_Call) RunAndReturn(run func() ([]Vet, error)) *MockVetRepositorier_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllPreload provides a mock function with given fields:
func (_m *MockVetRepositorier) FindAllPreload() ([]Vet, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAllPreload")
	}

	var r0 []Vet
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Vet, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Vet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Vet)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetRepositorier_FindAllPreload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllPreload'
type MockVetRepositorier_FindAllPreload_Call struct {
	*mock.Call
}

// FindAllPreload is a helper method to define mock.On call
func (_e *MockVetRepositorier_Expecter) FindAllPreload() *MockVetRepositorier_FindAllPreload_Call {
	return &MockVetRepositorier_FindAllPreload_Call{Call: _e.mock.On("FindAllPreload")}
}

func (_c *MockVetRepositorier_FindAllPreload_Call) Run(run func()) *MockVetRepositorier_FindAllPreload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVetRepositorier_FindAllPreload_Call) Return(_a0 []Vet, _a1 error) *MockVetRepositorier_FindAllPreload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetRepositorier_FindAllPreload_Call) RunAndReturn(run func() ([]Vet, error)) *MockVetRepositorier_FindAllPreload_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *MockVetRepositorier) FindById(id int) (*Vet, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *Vet
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*Vet, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *Vet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Vet)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetRepositorier_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockVetRepositorier_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id int
func (_e *MockVetRepositorier_Expecter) FindById(id interface{}) *MockVetRepositorier_FindById_Call {
	return &MockVetRepositorier_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *MockVetRepositorier_FindById_Call) Run(run func(id int)) *MockVetRepositorier_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockVetRepositorier_FindById_Call) Return(_a0 *Vet, _a1 error) *MockVetRepositorier_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetRepositorier_FindById_Call) RunAndReturn(run func(int) (*Vet, error)) *MockVetRepositorier_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// FindByLastName provides a mock function with given fields: lastName
func (_m *MockVetRepositorier) FindByLastName(lastName string) ([]Vet, error) {
	ret := _m.Called(lastName)

	if len(ret) == 0 {
		panic("no return value specified for FindByLastName")
	}

	var r0 []Vet
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]Vet, error)); ok {
		return rf(lastName)
	}
	if rf, ok := ret.Get(0).(func(string) []Vet); ok {
		r0 = rf(lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Vet)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetRepositorier_FindByLastName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByLastName'
type MockVetRepositorier_FindByLastName_Call struct {
	*mock.Call
}

// FindByLastName is a helper method to define mock.On call
//   - lastName string
func (_e *MockVetRepositorier_Expecter) FindByLastName(lastName interface{}) *MockVetRepositorier_FindByLastName_Call {
	return &MockVetRepositorier_FindByLastName_Call{Call: _e.mock.On("FindByLastName", lastName)}
}

func (_c *MockVetRepositorier_FindByLastName_Call) Run(run func(lastName string)) *MockVetRepositorier_FindByLastName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVetRepositorier_FindByLastName_Call) Return(_a0 []Vet, _a1 error) *MockVetRepositorier_FindByLastName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetRepositorier_FindByLastName_Call) RunAndReturn(run func(string) ([]Vet, error)) *MockVetRepositorier_FindByLastName_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: vet
func (_m *MockVetRepositorier) Insert(vet *Vet) (*Vet, error) {
	ret := _m.Called(vet)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 *Vet
	var r1 error
	if rf, ok := ret.Get(0).(func(*Vet) (*Vet, error)); ok {
		return rf(vet)
	}
	if rf, ok := ret.Get(0).(func(*Vet) *Vet); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Vet)
		}
	}

	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetRepositorier_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockVetRepositorier_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - vet *Vet
func (_e *MockVetRepositorier_Expecter) Insert(vet interface{}) *MockVetRepositorier_Insert_Call {
	return &MockVetRepositorier_Insert_Call{Call: _e.mock.On("Insert", vet)}
}

func (_c *MockVetRepositorier_Insert_Call) Run(run func(vet *Vet)) *MockVetRepositorier_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Vet))
	})
	return _c
}

func (_c *MockVetRepositorier_Insert_Call) Return(_a0 *Vet, _a1 error) *MockVetRepositorier_Insert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetRepositorier_Insert_Call) RunAndReturn(run func(*Vet) (*Vet, error)) *MockVetRepositorier_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: vet
func (_m *MockVetRepositorier) Update(vet *Vet) (*Vet, error) {
	ret := _m.Called(vet)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *Vet
	var r1 error
	if rf, ok := ret.Get(0).(func(*Vet) (*Vet, error)); ok {
		return rf(vet)
	}
	if rf, ok := ret.Get(0).(func(*Vet) *Vet); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Vet)
		}
	}

	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVetRepositorier_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockVetRepositorier_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - vet *Vet
func (_e *MockVetRepositorier_Expecter) Update(vet interface{}) *MockVetRepositorier_Update_Call {
	return &MockVetRepositorier_Update_Call{Call: _e.mock.On("Update", vet)}
}

func (_c *MockVetRepositorier_Update_Call) Run(run func(vet *Vet)) *MockVetRepositorier_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Vet))
	})
	return _c
}

func (_c *MockVetRepositorier_Update_Call) Return(_a0 *Vet, _a1 error) *MockVetRepositorier_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVetRepositorier_Update_Call) RunAndReturn(run func(*Vet) (*Vet, error)) *MockVetRepositorier_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVetRepositorier creates a new instance of MockVetRepositorier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVetRepositorier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVetRepositorier {
	mock := &MockVetRepositorier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
