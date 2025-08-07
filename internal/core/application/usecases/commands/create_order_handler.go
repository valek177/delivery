package commands

import (
	"context"

	"delivery/internal/core/domain/model/order"
	"delivery/internal/core/ports"
	"delivery/internal/pkg/errs"
)

type CreateOrderCommandHandler interface {
	Handle(context.Context, CreateOrderCommand) error
}

var _ CreateOrderCommandHandler = &createOrderCommandHandler{}

type createOrderCommandHandler struct {
	uowFactory ports.UnitOfWorkFactory
	geoClient  ports.GeoClient
}

func NewCreateOrderCommandHandler(
	uowFactory ports.UnitOfWorkFactory, geoClient ports.GeoClient,
) (CreateOrderCommandHandler, error) {
	if uowFactory == nil {
		return nil, errs.NewValueIsRequiredError("unitOfWork")
	}

	if geoClient == nil {
		return nil, errs.NewValueIsRequiredError("geoClient")
	}

	return &createOrderCommandHandler{
		uowFactory: uowFactory,
		geoClient:  geoClient,
	}, nil
}

func (ch *createOrderCommandHandler) Handle(ctx context.Context, command CreateOrderCommand) error {
	// Создаем UoW
	uow, err := ch.uowFactory.New(ctx)
	if err != nil {
		return err
	}
	defer uow.RollbackUnlessCommitted(ctx)

	// Проверяем нет ли уже такого заказа
	orderAggregate, err := uow.OrderRepository().Get(ctx, command.OrderID())
	if err != nil {
		return err
	}
	if orderAggregate != nil {
		return nil
	}

	// Получаем геопозицию из сервиса Geo.
	location, err := ch.geoClient.GetGeolocation(ctx, command.Street())
	if err != nil {
		return err
	}

	// Изменили
	orderAggregate, err = order.NewOrder(command.OrderID(), location, command.Volume())
	if err != nil {
		return err
	}

	// Сохранили
	err = uow.OrderRepository().Add(ctx, orderAggregate)
	if err != nil {
		return err
	}

	return nil
}
