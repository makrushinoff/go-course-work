// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	model "lab3/model"

	mock "github.com/stretchr/testify/mock"
)

// IBoardRepository is an autogenerated mock type for the IBoardRepository type
type IBoardRepository struct {
	mock.Mock
}

// CreateBoard provides a mock function with given fields: board
func (_m *IBoardRepository) CreateBoard(board model.BoardDto) error {
	ret := _m.Called(board)

	if len(ret) == 0 {
		panic("no return value specified for CreateBoard")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.BoardDto) error); ok {
		r0 = rf(board)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBoards provides a mock function with given fields:
func (_m *IBoardRepository) GetAllBoards() ([]model.BoardDto, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllBoards")
	}

	var r0 []model.BoardDto
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.BoardDto, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.BoardDto); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.BoardDto)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBoardById provides a mock function with given fields: id
func (_m *IBoardRepository) GetBoardById(id string) (model.BoardDto, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetBoardById")
	}

	var r0 model.BoardDto
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.BoardDto, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) model.BoardDto); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.BoardDto)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIBoardRepository creates a new instance of IBoardRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIBoardRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IBoardRepository {
	mock := &IBoardRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
