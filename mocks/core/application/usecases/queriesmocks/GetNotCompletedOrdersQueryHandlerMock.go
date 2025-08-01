// Code generated by mockery. DO NOT EDIT.

package queriesmocks

import (
	queries "delivery/internal/core/application/usecases/queries"

	mock "github.com/stretchr/testify/mock"
)

// GetNotCompletedOrdersQueryHandlerMock is an autogenerated mock type for the GetNotCompletedOrdersQueryHandler type
type GetNotCompletedOrdersQueryHandlerMock struct {
	mock.Mock
}

type GetNotCompletedOrdersQueryHandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *GetNotCompletedOrdersQueryHandlerMock) EXPECT() *GetNotCompletedOrdersQueryHandlerMock_Expecter {
	return &GetNotCompletedOrdersQueryHandlerMock_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: _a0
func (_m *GetNotCompletedOrdersQueryHandlerMock) Handle(_a0 queries.GetNotCompletedOrdersQuery) (queries.GetNotCompletedOrdersResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 queries.GetNotCompletedOrdersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(queries.GetNotCompletedOrdersQuery) (queries.GetNotCompletedOrdersResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(queries.GetNotCompletedOrdersQuery) queries.GetNotCompletedOrdersResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(queries.GetNotCompletedOrdersResponse)
	}

	if rf, ok := ret.Get(1).(func(queries.GetNotCompletedOrdersQuery) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotCompletedOrdersQueryHandlerMock_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type GetNotCompletedOrdersQueryHandlerMock_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - _a0 queries.GetNotCompletedOrdersQuery
func (_e *GetNotCompletedOrdersQueryHandlerMock_Expecter) Handle(_a0 interface{}) *GetNotCompletedOrdersQueryHandlerMock_Handle_Call {
	return &GetNotCompletedOrdersQueryHandlerMock_Handle_Call{Call: _e.mock.On("Handle", _a0)}
}

func (_c *GetNotCompletedOrdersQueryHandlerMock_Handle_Call) Run(run func(_a0 queries.GetNotCompletedOrdersQuery)) *GetNotCompletedOrdersQueryHandlerMock_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(queries.GetNotCompletedOrdersQuery))
	})
	return _c
}

func (_c *GetNotCompletedOrdersQueryHandlerMock_Handle_Call) Return(_a0 queries.GetNotCompletedOrdersResponse, _a1 error) *GetNotCompletedOrdersQueryHandlerMock_Handle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GetNotCompletedOrdersQueryHandlerMock_Handle_Call) RunAndReturn(run func(queries.GetNotCompletedOrdersQuery) (queries.GetNotCompletedOrdersResponse, error)) *GetNotCompletedOrdersQueryHandlerMock_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetNotCompletedOrdersQueryHandlerMock creates a new instance of GetNotCompletedOrdersQueryHandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetNotCompletedOrdersQueryHandlerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetNotCompletedOrdersQueryHandlerMock {
	mock := &GetNotCompletedOrdersQueryHandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
