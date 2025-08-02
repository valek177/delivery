package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"delivery/internal/adapters/in/http/problems"
	"delivery/internal/core/application/usecases/queries"
	"delivery/internal/generated/servers"
	"delivery/internal/pkg/errs"
)

func (s *Server) GetCouriers(c echo.Context) error {
	query, err := queries.NewGetCouriersQuery()
	if err != nil {
		return problems.NewBadRequest(err.Error())
	}

	queryResponse, err := s.getCouriersQueryHandler.Handle(query)
	if err != nil {
		if errors.Is(err, errs.ErrObjectNotFound) {
			return c.JSON(http.StatusNotFound, problems.NewNotFound(err.Error()))
		}
	}

	httpResponse := make([]servers.Courier, 0, len(queryResponse.Couriers))
	for _, courier := range queryResponse.Couriers {
		courier := servers.Courier{
			Id:   courier.ID,
			Name: courier.Name,
			Location: servers.Location{
				X: courier.Location.X,
				Y: courier.Location.Y,
			},
		}
		httpResponse = append(httpResponse, courier)
	}
	return c.JSON(http.StatusOK, httpResponse)
}
