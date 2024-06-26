// Code generated by mockery v2.42.2. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockPetRepositorier is an autogenerated mock type for the PetRepositorier type
type MockPetRepositorier struct {
	mock.Mock
}

type MockPetRepositorier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPetRepositorier) EXPECT() *MockPetRepositorier_Expecter {
	return &MockPetRepositorier_Expecter{mock: &_m.Mock}
}

// FindById provides a mock function with given fields: id
func (_m *MockPetRepositorier) FindById(id int) (*Pet, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *Pet
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*Pet, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *Pet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Pet)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPetRepositorier_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockPetRepositorier_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id int
func (_e *MockPetRepositorier_Expecter) FindById(id interface{}) *MockPetRepositorier_FindById_Call {
	return &MockPetRepositorier_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *MockPetRepositorier_FindById_Call) Run(run func(id int)) *MockPetRepositorier_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockPetRepositorier_FindById_Call) Return(_a0 *Pet, _a1 error) *MockPetRepositorier_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPetRepositorier_FindById_Call) RunAndReturn(run func(int) (*Pet, error)) *MockPetRepositorier_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// FindByName provides a mock function with given fields: name
func (_m *MockPetRepositorier) FindByName(name string) ([]Pet, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for FindByName")
	}

	var r0 []Pet
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]Pet, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) []Pet); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Pet)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPetRepositorier_FindByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByName'
type MockPetRepositorier_FindByName_Call struct {
	*mock.Call
}

// FindByName is a helper method to define mock.On call
//   - name string
func (_e *MockPetRepositorier_Expecter) FindByName(name interface{}) *MockPetRepositorier_FindByName_Call {
	return &MockPetRepositorier_FindByName_Call{Call: _e.mock.On("FindByName", name)}
}

func (_c *MockPetRepositorier_FindByName_Call) Run(run func(name string)) *MockPetRepositorier_FindByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPetRepositorier_FindByName_Call) Return(_a0 []Pet, _a1 error) *MockPetRepositorier_FindByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPetRepositorier_FindByName_Call) RunAndReturn(run func(string) ([]Pet, error)) *MockPetRepositorier_FindByName_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: pet
func (_m *MockPetRepositorier) Insert(pet *Pet) (*Pet, error) {
	ret := _m.Called(pet)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 *Pet
	var r1 error
	if rf, ok := ret.Get(0).(func(*Pet) (*Pet, error)); ok {
		return rf(pet)
	}
	if rf, ok := ret.Get(0).(func(*Pet) *Pet); ok {
		r0 = rf(pet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Pet)
		}
	}

	if rf, ok := ret.Get(1).(func(*Pet) error); ok {
		r1 = rf(pet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPetRepositorier_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockPetRepositorier_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - pet *Pet
func (_e *MockPetRepositorier_Expecter) Insert(pet interface{}) *MockPetRepositorier_Insert_Call {
	return &MockPetRepositorier_Insert_Call{Call: _e.mock.On("Insert", pet)}
}

func (_c *MockPetRepositorier_Insert_Call) Run(run func(pet *Pet)) *MockPetRepositorier_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Pet))
	})
	return _c
}

func (_c *MockPetRepositorier_Insert_Call) Return(_a0 *Pet, _a1 error) *MockPetRepositorier_Insert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPetRepositorier_Insert_Call) RunAndReturn(run func(*Pet) (*Pet, error)) *MockPetRepositorier_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: pet
func (_m *MockPetRepositorier) Update(pet *Pet) (*Pet, error) {
	ret := _m.Called(pet)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *Pet
	var r1 error
	if rf, ok := ret.Get(0).(func(*Pet) (*Pet, error)); ok {
		return rf(pet)
	}
	if rf, ok := ret.Get(0).(func(*Pet) *Pet); ok {
		r0 = rf(pet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Pet)
		}
	}

	if rf, ok := ret.Get(1).(func(*Pet) error); ok {
		r1 = rf(pet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPetRepositorier_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPetRepositorier_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - pet *Pet
func (_e *MockPetRepositorier_Expecter) Update(pet interface{}) *MockPetRepositorier_Update_Call {
	return &MockPetRepositorier_Update_Call{Call: _e.mock.On("Update", pet)}
}

func (_c *MockPetRepositorier_Update_Call) Run(run func(pet *Pet)) *MockPetRepositorier_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Pet))
	})
	return _c
}

func (_c *MockPetRepositorier_Update_Call) Return(_a0 *Pet, _a1 error) *MockPetRepositorier_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPetRepositorier_Update_Call) RunAndReturn(run func(*Pet) (*Pet, error)) *MockPetRepositorier_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPetRepositorier creates a new instance of MockPetRepositorier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPetRepositorier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPetRepositorier {
	mock := &MockPetRepositorier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
