package http

import (
	"delivery/internal/core/application/usecases/commands"
	"delivery/internal/core/application/usecases/queries"
	"delivery/internal/pkg/errs"
)

type Server struct {
	createOrderCommandHandler   commands.CreateOrderCommandHandler
	createCourierCommandHandler commands.CreateCourierCommandHandler

	getCouriersQueryHandler           queries.GetCouriersQueryHandler
	getNotCompletedOrdersQueryHandler queries.GetNotCompletedOrdersQueryHandler
}

func NewServer(
	createOrderCommandHandler commands.CreateOrderCommandHandler,
	createCourierCommandHandler commands.CreateCourierCommandHandler,

	getCouriersQueryHandler queries.GetCouriersQueryHandler,
	getNotCompletedOrdersQueryHandler queries.GetNotCompletedOrdersQueryHandler,
) (*Server, error) {
	if createOrderCommandHandler == nil {
		return nil, errs.NewValueIsRequiredError("createOrderCommandHandler")
	}
	if createCourierCommandHandler == nil {
		return nil, errs.NewValueIsRequiredError("createCourierCommandHandler")
	}
	if getCouriersQueryHandler == nil {
		return nil, errs.NewValueIsRequiredError("getCouriersQueryHandler")
	}
	if getNotCompletedOrdersQueryHandler == nil {
		return nil, errs.NewValueIsRequiredError("getNotCompletedOrdersQueryHandler")
	}
	return &Server{
		createOrderCommandHandler:         createOrderCommandHandler,
		createCourierCommandHandler:       createCourierCommandHandler,
		getCouriersQueryHandler:           getCouriersQueryHandler,
		getNotCompletedOrdersQueryHandler: getNotCompletedOrdersQueryHandler,
	}, nil
}
