// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	models "github.com/irahardianto/service-pattern-go/models"
	mock "github.com/stretchr/testify/mock"
)

// IPlayerRepository is an autogenerated mock type for the IPlayerRepository type
type IPlayerRepository struct {
	mock.Mock
}

// GetPlayerByName provides a mock function with given fields: name
func (_m *IPlayerRepository) GetPlayerByName(name string) (*models.Player, error) {
	ret := _m.Called(name)

	var r0 *models.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Player, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Player); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Player)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIPlayerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPlayerRepository creates a new instance of IPlayerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPlayerRepository(t mockConstructorTestingTNewIPlayerRepository) *IPlayerRepository {
	mock := &IPlayerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
