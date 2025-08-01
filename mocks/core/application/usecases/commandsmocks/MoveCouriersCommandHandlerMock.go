// Code generated by mockery. DO NOT EDIT.

package commandsmocks

import (
	context "context"
	commands "delivery/internal/core/application/usecases/commands"

	mock "github.com/stretchr/testify/mock"
)

// MoveCouriersCommandHandlerMock is an autogenerated mock type for the MoveCouriersCommandHandler type
type MoveCouriersCommandHandlerMock struct {
	mock.Mock
}

type MoveCouriersCommandHandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *MoveCouriersCommandHandlerMock) EXPECT() *MoveCouriersCommandHandlerMock_Expecter {
	return &MoveCouriersCommandHandlerMock_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: _a0, _a1
func (_m *MoveCouriersCommandHandlerMock) Handle(_a0 context.Context, _a1 commands.MoveCouriersCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.MoveCouriersCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveCouriersCommandHandlerMock_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MoveCouriersCommandHandlerMock_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 commands.MoveCouriersCommand
func (_e *MoveCouriersCommandHandlerMock_Expecter) Handle(_a0 interface{}, _a1 interface{}) *MoveCouriersCommandHandlerMock_Handle_Call {
	return &MoveCouriersCommandHandlerMock_Handle_Call{Call: _e.mock.On("Handle", _a0, _a1)}
}

func (_c *MoveCouriersCommandHandlerMock_Handle_Call) Run(run func(_a0 context.Context, _a1 commands.MoveCouriersCommand)) *MoveCouriersCommandHandlerMock_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(commands.MoveCouriersCommand))
	})
	return _c
}

func (_c *MoveCouriersCommandHandlerMock_Handle_Call) Return(_a0 error) *MoveCouriersCommandHandlerMock_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MoveCouriersCommandHandlerMock_Handle_Call) RunAndReturn(run func(context.Context, commands.MoveCouriersCommand) error) *MoveCouriersCommandHandlerMock_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMoveCouriersCommandHandlerMock creates a new instance of MoveCouriersCommandHandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMoveCouriersCommandHandlerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *MoveCouriersCommandHandlerMock {
	mock := &MoveCouriersCommandHandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
