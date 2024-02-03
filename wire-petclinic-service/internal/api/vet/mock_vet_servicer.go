// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package vet

import mock "github.com/stretchr/testify/mock"

// MockVetServicer is an autogenerated mock type for the VetServicer type
type MockVetServicer struct {
	mock.Mock
}

// Create provides a mock function with given fields: vet
func (_m *MockVetServicer) Create(vet *Vet) (*VetResponse, error) {
	ret := _m.Called(vet)

	var r0 *VetResponse
	if rf, ok := ret.Get(0).(func(*Vet) *VetResponse); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllVets provides a mock function with given fields:
func (_m *MockVetServicer) GetAllVets() ([]VetResponse, error) {
	ret := _m.Called()

	var r0 []VetResponse
	if rf, ok := ret.Get(0).(func() []VetResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]VetResponse)
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

// GetVetById provides a mock function with given fields: id
func (_m *MockVetServicer) GetVetById(id int) (*VetResponse, error) {
	ret := _m.Called(id)

	var r0 *VetResponse
	if rf, ok := ret.Get(0).(func(int) *VetResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VetResponse)
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

// GetVetByLastName provides a mock function with given fields: lastName
func (_m *MockVetServicer) GetVetByLastName(lastName string) ([]VetResponse, error) {
	ret := _m.Called(lastName)

	var r0 []VetResponse
	if rf, ok := ret.Get(0).(func(string) []VetResponse); ok {
		r0 = rf(lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]VetResponse)
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

// Update provides a mock function with given fields: vet
func (_m *MockVetServicer) Update(vet *Vet) (*VetResponse, error) {
	ret := _m.Called(vet)

	var r0 *VetResponse
	if rf, ok := ret.Get(0).(func(*Vet) *VetResponse); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}