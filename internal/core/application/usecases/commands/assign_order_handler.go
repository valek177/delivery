package commands

import (
	"context"
	"errors"

	"delivery/internal/core/domain/services"
	"delivery/internal/core/ports"
	"delivery/internal/pkg/errs"
)

type AssignOrderCommandHandler interface {
	Handle(context.Context, AssignOrderCommand) error
}

var _ AssignOrderCommandHandler = &assignOrderCommandHandler{}

type assignOrderCommandHandler struct {
	uowFactory      ports.UnitOfWorkFactory
	orderDispatcher services.OrderDispatcher
}

func NewAssignOrderCommandHandler(
	uowFactory ports.UnitOfWorkFactory, orderDispatcher services.OrderDispatcher,
) (AssignOrderCommandHandler, error) {
	if uowFactory == nil {
		return nil, errs.NewValueIsRequiredError("unitOfWork")
	}
	if orderDispatcher == nil {
		return nil, errs.NewValueIsRequiredError("orderDispatcher")
	}

	return &assignOrderCommandHandler{
		uowFactory:      uowFactory,
		orderDispatcher: orderDispatcher,
	}, nil
}

func (ch *assignOrderCommandHandler) Handle(ctx context.Context, command AssignOrderCommand) error {
	uow, err := ch.uowFactory.New(ctx)
	if err != nil {
		return err
	}
	defer uow.RollbackUnlessCommitted(ctx)

	createdOrder, err := uow.OrderRepository().GetFirstInCreatedStatus(ctx)
	if err != nil {
		if errors.Is(err, errs.ErrObjectNotFound) {
			return errors.New("unable to get orders")
		}
		return err
	}

	freeCouriers, err := uow.CourierRepository().GetAllFree(ctx)
	if err != nil {
		if errors.Is(err, errs.ErrObjectNotFound) {
			return errors.New("unable to get couriers")
		}
		return err
	}

	if len(freeCouriers) == 0 {
		return errors.New("no free couriers")
	}

	courier, err := ch.orderDispatcher.Dispatch(createdOrder, freeCouriers)
	if err != nil {
		return err
	}

	uow.Begin(ctx)
	err = uow.OrderRepository().Update(ctx, createdOrder)
	if err != nil {
		return err
	}

	err = uow.CourierRepository().Update(ctx, courier)
	if err != nil {
		return err
	}

	err = uow.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
