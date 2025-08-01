package commands

import (
	"context"

	"delivery/internal/core/domain/kernel"
	"delivery/internal/core/domain/model/courier"
	"delivery/internal/core/ports"
	"delivery/internal/pkg/errs"
)

type CreateCourierCommandHandler interface {
	Handle(context.Context, CreateCourierCommand) error
}

var _ CreateCourierCommandHandler = &createCourierCommandHandler{}

type createCourierCommandHandler struct {
	uowFactory ports.UnitOfWorkFactory
}

func NewCreateCourierCommandHandler(
	uowFactory ports.UnitOfWorkFactory,
) (CreateCourierCommandHandler, error) {
	if uowFactory == nil {
		return nil, errs.NewValueIsRequiredError("unitOfWork")
	}

	return &createCourierCommandHandler{
		uowFactory: uowFactory,
	}, nil
}

func (ch *createCourierCommandHandler) Handle(ctx context.Context, command CreateCourierCommand) error {
	// Создаем UoW
	uow, err := ch.uowFactory.New(ctx)
	if err != nil {
		return err
	}
	defer uow.RollbackUnlessCommitted(ctx)

	// Создаем нового курьера
	location := kernel.NewRandomLocation()
	courierAggregate, err := courier.NewCourier(command.Name(), command.Speed(), location)
	if err != nil {
		return err
	}

	// Сохранили
	err = uow.CourierRepository().Add(ctx, courierAggregate)
	if err != nil {
		return err
	}

	return nil
}
