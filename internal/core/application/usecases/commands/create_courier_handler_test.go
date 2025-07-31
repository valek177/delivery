package commands

import (
	"context"
	"testing"

	"delivery/internal/core/domain/kernel"
	"delivery/internal/core/domain/model/courier"
	"delivery/mocks/core/portsmocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateCourierCommandHandlerShouldBeSuccessWhenParamsAreCorrect(t *testing.T) {
	// Arrange
	ctx := context.Background()
	location := kernel.MinLocation()
	courierAggregate, err := courier.NewCourier("Courier", 3, location)
	assert.NoError(t, err)

	var capturedObj *courier.Courier
	courierRepositoryMock := &portsmocks.CourierRepositoryMock{}
	courierRepositoryMock.
		On("Get", ctx, courierAggregate.ID()).
		Return(nil, nil).
		Once()
	courierRepositoryMock.
		On("Add", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			capturedObj = args.Get(1).(*courier.Courier)
		}).
		Return(nil, nil).
		Once()

	unitOfWorkMock := &portsmocks.UnitOfWorkMock{}
	unitOfWorkMock.
		On("CourierRepository").
		Return(courierRepositoryMock)
	unitOfWorkFactoryMock := &portsmocks.UnitOfWorkFactoryMock{}
	unitOfWorkFactoryMock.
		On("New", ctx).
		Return(unitOfWorkMock, nil)
	unitOfWorkMock.
		On("RollbackUnlessCommitted", ctx).
		Return()

	// Act
	createCourierCommandHandler, err := NewCreateCourierCommandHandler(unitOfWorkFactoryMock)
	assert.NoError(t, err)
	createCourierCommand, err := NewCreateCourierCommand(courierAggregate.Name(), 3)
	assert.NoError(t, err)
	err = createCourierCommandHandler.Handle(ctx, createCourierCommand)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, capturedObj.ID())
}
