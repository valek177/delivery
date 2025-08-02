package http

import (
	"errors"
	"net/http"

	"delivery/internal/adapters/in/http/problems"
	"delivery/internal/core/application/usecases/queries"
	"delivery/internal/generated/servers"
	"delivery/internal/pkg/errs"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetOrders(c echo.Context) error {
	query, err := queries.NewGetNotCompletedOrdersQuery()
	if err != nil {
		return problems.NewBadRequest(err.Error())
	}

	queryResponse, err := s.getNotCompletedOrdersQueryHandler.Handle(query)
	if err != nil {
		if errors.Is(err, errs.ErrObjectNotFound) {
			return c.JSON(http.StatusNotFound, problems.NewNotFound(err.Error()))
		}
	}

	httpResponse := make([]servers.Order, 0, len(queryResponse.Orders))
	for _, courier := range queryResponse.Orders {
		courier := servers.Order{
			Id: courier.ID,
			Location: servers.Location{
				X: courier.Location.X,
				Y: courier.Location.Y,
			},
		}

		httpResponse = append(httpResponse, courier)
	}
	return c.JSON(http.StatusOK, httpResponse)
}
