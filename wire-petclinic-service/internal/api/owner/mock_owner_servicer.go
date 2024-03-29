// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package owner

import mock "github.com/stretchr/testify/mock"

// MockOwnerServicer is an autogenerated mock type for the Servicer type
type MockOwnerServicer struct {
	mock.Mock
}

// Create provides a mock function with given fields: owner
func (_m *MockOwnerServicer) Create(owner *Owner) (*OwnerResponse, error) {
	ret := _m.Called(owner)

	var r0 *OwnerResponse
	if rf, ok := ret.Get(0).(func(*Owner) *OwnerResponse); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*OwnerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Owner) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOwners provides a mock function with given fields:
func (_m *MockOwnerServicer) GetAllOwners() ([]OwnerResponse, error) {
	ret := _m.Called()

	var r0 []OwnerResponse
	if rf, ok := ret.Get(0).(func() []OwnerResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]OwnerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOwnersWithPets provides a mock function with given fields:
func (_m *MockOwnerServicer) GetAllOwnersWithPets() ([]OwnerResponse, error) {
	ret := _m.Called()

	var r0 []OwnerResponse
	if rf, ok := ret.Get(0).(func() []OwnerResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]OwnerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOwnerById provides a mock function with given fields: id
func (_m *MockOwnerServicer) GetOwnerById(id int) (*OwnerResponse, error) {
	ret := _m.Called(id)

	var r0 *OwnerResponse
	if rf, ok := ret.Get(0).(func(int) *OwnerResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*OwnerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOwnerByLastName provides a mock function with given fields: lastName
func (_m *MockOwnerServicer) GetOwnerByLastName(lastName string) ([]OwnerResponse, error) {
	ret := _m.Called(lastName)

	var r0 []OwnerResponse
	if rf, ok := ret.Get(0).(func(string) []OwnerResponse); ok {
		r0 = rf(lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]OwnerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: owner
func (_m *MockOwnerServicer) Update(owner *Owner) (*OwnerResponse, error) {
	ret := _m.Called(owner)

	var r0 *OwnerResponse
	if rf, ok := ret.Get(0).(func(*Owner) *OwnerResponse); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*OwnerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Owner) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
